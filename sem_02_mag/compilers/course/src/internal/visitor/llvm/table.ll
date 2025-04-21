%TableEntry = type {
  %Generic*,        ; ключ
  %Generic*,        ; значение
  i1                ; флаг занятости
}

%LuaTable = type {
  i32,              ; текущий размер
  i32,              ; вместимость
  %TableEntry*      ; массив записей
}

@TABLE_TYPE = constant i32 4
@INITIAL_CAPACITY = constant i32 16

; =============================================
; Исправленная функция создания таблицы
; =============================================
define %Generic* @lua_table_new() {
entry:
  %capacity = load i32, i32* @INITIAL_CAPACITY
  
  ; Вычисление размера одного элемента TableEntry
  %entry_size = ptrtoint %TableEntry* getelementptr inbounds (%TableEntry, %TableEntry* null, i32 1) to i64
  
  ; Преобразуем capacity в i64
  %capacity_sext = sext i32 %capacity to i64
  
  ; Вычисление общего размера массива
  %entries_size = mul i64 %entry_size, %capacity_sext  ; Используем capacity_sext
  
  ; Выделение памяти для массива записей
  %entries = call i8* @malloc(i64 %entries_size)
  %entries_ptr = bitcast i8* %entries to %TableEntry*
  
  ; Инициализация массива
  br label %init_loop

init_loop:
  %i = phi i32 [0, %entry], [%next_i, %init_loop]
  %entry_ptr = getelementptr inbounds %TableEntry, %TableEntry* %entries_ptr, i32 %i
  %flag_ptr = getelementptr inbounds %TableEntry, %TableEntry* %entry_ptr, i32 0, i32 2
  store i1 false, i1* %flag_ptr, align 1
  %next_i = add i32 %i, 1
  %done = icmp eq i32 %next_i, %capacity
  br i1 %done, label %create_table, label %init_loop

create_table:
  ; Выделение памяти для LuaTable
  %table_size = ptrtoint %LuaTable* getelementptr inbounds (%LuaTable, %LuaTable* null, i32 1) to i64
  %table = call i8* @malloc(i64 %table_size)
  %null_table = icmp eq i8* %table, null
  br i1 %null_table, label %error, label %continue

continue:
  %table_ptr = bitcast i8* %table to %LuaTable*
  
  ; Инициализация структуры LuaTable
  %size_ptr = getelementptr inbounds %LuaTable, %LuaTable* %table_ptr, i32 0, i32 0
  store i32 0, i32* %size_ptr, align 4
  %cap_ptr = getelementptr inbounds %LuaTable, %LuaTable* %table_ptr, i32 0, i32 1
  store i32 %capacity, i32* %cap_ptr, align 4
  %entries_field = getelementptr inbounds %LuaTable, %LuaTable* %table_ptr, i32 0, i32 2
  store %TableEntry* %entries_ptr, %TableEntry** %entries_field, align 8
  
  ; Упаковка в Generic с типом 4 (таблица)
  %generic = call %Generic* @create(i32 4, i8* %table)
  ret %Generic* %generic

error:
  call void @panic(i8* getelementptr inbounds ([11 x i8], [11 x i8]* @.error.null_table, i32 0, i32 0))
  ret %Generic* null
}

; =============================================
; Исправленная функция сохранения значения
; =============================================
define void @lua_table_set(%Generic* %table, %Generic* %key, %Generic* %value) {
entry:
  %tbl = call %LuaTable* @extract_table(%Generic* %table)
  %is_valid = icmp ne %LuaTable* %tbl, null
  br i1 %is_valid, label %proceed, label %error

proceed:
  ; Копирование ключа и значения
  %key_copy = call %Generic* @create_nil()
  call void @copy(%Generic* %key, %Generic* %key_copy)
  
  %value_copy = call %Generic* @create_nil()
  call void @copy(%Generic* %value, %Generic* %value_copy)

  ; Получение данных таблицы
  %entries_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 2
  %entries = load %TableEntry*, %TableEntry** %entries_ptr, align 8
  %capacity_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 1
  %capacity = load i32, i32* %capacity_ptr, align 4
  %size_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 0
  %size = load i32, i32* %size_ptr, align 4
  
  br label %search_loop

search_loop:
  %i = phi i32 [0, %proceed], [%next_i, %next]
  %current = getelementptr inbounds %TableEntry, %TableEntry* %entries, i32 %i
  %occupied_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current, i32 0, i32 2
  %occupied = load i1, i1* %occupied_ptr, align 1
  br i1 %occupied, label %check_key, label %try_insert

check_key:
  %current_key_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current, i32 0, i32 0
  %current_key = load %Generic*, %Generic** %current_key_ptr, align 8
  %is_equal = call %Generic* @equal(%Generic* %current_key, %Generic* %key)
  %is_true = call i1 @check(%Generic* %is_equal)
  br i1 %is_true, label %update, label %next

update:
  %old_value_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current, i32 0, i32 1
  %old_value = load %Generic*, %Generic** %old_value_ptr, align 8
  call void @destroy(%Generic* %old_value)
  store %Generic* %value_copy, %Generic** %old_value_ptr, align 8
  br label %exit

try_insert:
  %key_slot_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current, i32 0, i32 0
  store %Generic* %key_copy, %Generic** %key_slot_ptr, align 8
  %value_slot_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current, i32 0, i32 1
  store %Generic* %value_copy, %Generic** %value_slot_ptr, align 8
  store i1 true, i1* %occupied_ptr, align 1
  %new_size = add i32 %size, 1
  store i32 %new_size, i32* %size_ptr, align 4
  br label %exit

next:
  %next_i = add i32 %i, 1
  %in_bounds = icmp ult i32 %next_i, %capacity
  br i1 %in_bounds, label %search_loop, label %resize

resize:
  call void @resize_table(%LuaTable* %tbl)
  call void @lua_table_set(%Generic* %table, %Generic* %key, %Generic* %value)
  br label %exit

exit:
  ret void

error:
  call void @panic(i8* getelementptr inbounds ([21 x i8], [21 x i8]* @.error.cannot_extract_table, i32 0, i32 0))
  ret void
}

define %Generic* @lua_table_get(%Generic* %table, %Generic* %key) {
entry:
  %tbl = call %LuaTable* @extract_table(%Generic* %table)
  %valid = icmp ne %LuaTable* %tbl, null
  br i1 %valid, label %proceed, label %error

proceed:
  %entries_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 2
  %entries = load %TableEntry*, %TableEntry** %entries_ptr, align 8
  %capacity_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 1
  %capacity = load i32, i32* %capacity_ptr, align 4
  
  br label %search_loop

search_loop:
  %i = phi i32 [0, %proceed], [%next_i, %next]
  %current = getelementptr inbounds %TableEntry, %TableEntry* %entries, i32 %i
  %occupied_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current, i32 0, i32 2
  %occupied = load i1, i1* %occupied_ptr, align 1
  br i1 %occupied, label %check_key, label %next

check_key:
  %current_key_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current, i32 0, i32 0
  %current_key = load %Generic*, %Generic** %current_key_ptr, align 8
  %is_equal = call %Generic* @equal(%Generic* %current_key, %Generic* %key)
  %is_true = call i1 @check(%Generic* %is_equal)
  br i1 %is_true, label %found, label %next

found:
  %value_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current, i32 0, i32 1
  %value = load %Generic*, %Generic** %value_ptr, align 8
  %value_copy = call %Generic* @create_nil()
  call void @copy(%Generic* %value, %Generic* %value_copy)
  ret %Generic* %value_copy

next:
  %next_i = add i32 %i, 1
  %in_bounds = icmp ult i32 %next_i, %capacity
  br i1 %in_bounds, label %search_loop, label %not_found

not_found:
  %nil = call %Generic* @create_nil()
  ret %Generic* %nil

error:
  call void @panic(i8* getelementptr inbounds ([21 x i8], [21 x i8]* @.error.cannot_extract_table, i32 0, i32 0))
  ret %Generic* null
}

define %LuaTable* @extract_table(%Generic* %table_generic) {
entry:
  %type_ptr = getelementptr inbounds %Generic, %Generic* %table_generic, i32 0, i32 0
  %type = load i32, i32* %type_ptr, align 4
  %is_table = icmp eq i32 %type, 4
  br i1 %is_table, label %valid, label %error

valid:
  %data_ptr = getelementptr inbounds %Generic, %Generic* %table_generic, i32 0, i32 1
  %data = load i8*, i8** %data_ptr, align 8
  %table = bitcast i8* %data to %LuaTable*
  ret %LuaTable* %table

error:
  call void @panic(i8* getelementptr inbounds ([22 x i8], [22 x i8]* @.error.table_expected, i32 0, i32 0))
  ret %LuaTable* null
}

; Строка для ошибки
@.error.table_expected = private constant [22 x i8] c"Value should be table\00"
@.error.cannot_extract_table = private constant [21 x i8] c"Cannot extract table\00"
@.error.null_table = private constant [11 x i8] c"Null table\00"

define void @resize_table(%LuaTable* %tbl) {
entry:
  %old_capacity_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 1
  %old_capacity = load i32, i32* %old_capacity_ptr, align 4
  %new_capacity = mul i32 %old_capacity, 2
  
  ; Выделение нового массива
  %entry_size = ptrtoint %TableEntry* getelementptr inbounds (%TableEntry, %TableEntry* null, i32 1) to i64
  %new_capacity_sext = sext i32 %new_capacity to i64
  %new_entries_size = mul i64 %entry_size, %new_capacity_sext
  %new_entries = call i8* @malloc(i64 %new_entries_size)
  %new_entries_ptr = bitcast i8* %new_entries to %TableEntry*
  
  ; Копирование данных из старого массива
  %old_entries_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 2
  %old_entries = load %TableEntry*, %TableEntry** %old_entries_ptr, align 8
  
  br label %copy_loop

copy_loop:
  ; Исправлено: %next вместо несуществующего %copy_step
  %i = phi i32 [0, %entry], [%next_i, %next]
  %current_old = getelementptr inbounds %TableEntry, %TableEntry* %old_entries, i32 %i
  %current_new = getelementptr inbounds %TableEntry, %TableEntry* %new_entries_ptr, i32 %i
  
  ; Копирование флага занятости
  %flag_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current_old, i32 0, i32 2
  %flag = load i1, i1* %flag_ptr, align 1
  %new_flag_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current_new, i32 0, i32 2
  store i1 %flag, i1* %new_flag_ptr, align 1
  
  br i1 %flag, label %copy_data, label %next

copy_data:
  ; Копирование ключа
  %old_key_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current_old, i32 0, i32 0
  %old_key = load %Generic*, %Generic** %old_key_ptr, align 8
  %new_key_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current_new, i32 0, i32 0
  %key_copy = call %Generic* @create_nil()
  call void @copy(%Generic* %old_key, %Generic* %key_copy)
  store %Generic* %key_copy, %Generic** %new_key_ptr, align 8
  
  ; Копирование значения
  %old_value_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current_old, i32 0, i32 1
  %old_value = load %Generic*, %Generic** %old_value_ptr, align 8
  %new_value_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current_new, i32 0, i32 1
  %value_copy = call %Generic* @create_nil()
  call void @copy(%Generic* %old_value, %Generic* %value_copy)
  store %Generic* %value_copy, %Generic** %new_value_ptr, align 8
  
  br label %next

next:
  %next_i = add i32 %i, 1
  %done = icmp eq i32 %next_i, %old_capacity
  br i1 %done, label %finish, label %copy_loop

finish:
  ; Освобождение старого массива
  %old_entries_i8 = bitcast %TableEntry* %old_entries to i8*
  call void @free(i8* %old_entries_i8)
  
  ; Обновление структуры таблицы
  %new_entries_field = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 2
  store %TableEntry* %new_entries_ptr, %TableEntry** %new_entries_field, align 8
  
  %new_capacity_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 1
  store i32 %new_capacity, i32* %new_capacity_ptr, align 4
  
  ret void
}