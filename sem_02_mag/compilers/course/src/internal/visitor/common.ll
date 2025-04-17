; Внешние объявления
declare i8* @malloc(i64)
declare void @free(i8*)
declare i32 @strcmp(i8*, i8*)
declare i64 @strlen(i8*)
declare i8* @strcpy(i8*, i8*)
declare i8* @strcat(i8*, i8*)
declare void @llvm.memcpy.p0i8.p0i8.i64(i8*, i8*, i64, i1 immarg)
declare i32 @printf(i8*, ...)
declare double @pow(double, double)

; Глобальные строки для форматирования вывода
@.str.int = private unnamed_addr constant [6 x i8] c"%lld\0A\00", align 1
@.str.float = private unnamed_addr constant [4 x i8] c"%f\0A\00", align 1
@.str.string = private unnamed_addr constant [4 x i8] c"%s\0A\00", align 1
@.str.true = private unnamed_addr constant [6 x i8] c"true\0A\00", align 1
@.str.false = private unnamed_addr constant [7 x i8] c"false\0A\00", align 1
@.str.nil = private unnamed_addr constant [5 x i8] c"nil\0A\00", align 1

; Определение обобщённой структуры
%Generic = type {
  i32,     ; тип данных (0=int, 1=float, 2=string, 3=bool)
  i8*      ; указатель на данные
}

define %Generic* @create_nil() {
  ret %Generic* null
}

define %Generic* @create(i32 %type, i8* %value) {
entry:
  %obj = call i8* @malloc(i64 16)
  %g = bitcast i8* %obj to %Generic*
  %type_ptr = getelementptr inbounds %Generic, %Generic* %g, i32 0, i32 0
  store i32 %type, i32* %type_ptr
  
  switch i32 %type, label %invalid [
    i32 0, label %init_int
    i32 1, label %init_float
    i32 2, label %init_str
    i32 3, label %init_bool
  ]

init_int:
  %int_space = call i8* @malloc(i64 8)
  %int_ptr = bitcast i8* %int_space to i64*
  %int_val = ptrtoint i8* %value to i64
  store i64 %int_val, i64* %int_ptr
  %data_int = getelementptr inbounds %Generic, %Generic* %g, i32 0, i32 1
  store i8* %int_space, i8** %data_int
  ret %Generic* %g

init_float:
  %float_space = call i8* @malloc(i64 8)
  %float_ptr = bitcast i8* %float_space to double*
  %value_f = bitcast i8* %value to double*
  %f_val = load double, double* %value_f
  store double %f_val, double* %float_ptr
  %data_float = getelementptr inbounds %Generic, %Generic* %g, i32 0, i32 1
  store i8* %float_space, i8** %data_float
  ret %Generic* %g

init_str:
  %len = call i64 @strlen(i8* %value)
  %str_space = call i8* @malloc(i64 %len)
  call i8* @strcpy(i8* %str_space, i8* %value)
  %data_str = getelementptr inbounds %Generic, %Generic* %g, i32 0, i32 1
  store i8* %str_space, i8** %data_str
  ret %Generic* %g

init_bool:
  %bool_space = call i8* @malloc(i64 1)
  %bool_ptr = bitcast i8* %bool_space to i8*
  %bool_val = ptrtoint i8* %value to i8
  store i8 %bool_val, i8* %bool_ptr
  %data_bool = getelementptr inbounds %Generic, %Generic* %g, i32 0, i32 1
  store i8* %bool_space, i8** %data_bool
  ret %Generic* %g

invalid:
  ret %Generic* %g
}

define %Generic* @neg(%Generic* %v) {
entry:
  ; ИСПРАВЛЕННЫЙ ДОСТУП К ПОЛЯМ СТРУКТУРЫ
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %v, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  ; Проверка типов
  %type_eq_int = icmp eq i32 %a_type, 0
  br i1 %type_eq_int, label %neg_int, label %neg_float

neg_float:
  %v_fdata_ptr = getelementptr inbounds %Generic, %Generic* %v, i32 0, i32 1
  %v_fdata = load i8*, i8** %v_fdata_ptr
  %v_fptr = bitcast i8* %v_fdata to double*
  %v_fval = load double, double* %v_fptr
  %v_neg_fval = fneg double %v_fval
  
  ; Исправленное преобразование через временное хранение
  %temp.storage.neg = alloca double
  store double %v_neg_fval, double* %temp.storage.neg
  %as.i8.neg = bitcast double* %temp.storage.neg to i8*
  
  %fresult.neg = call %Generic* @create(i32 1, i8* %as.i8.neg)
  ret %Generic* %fresult.neg

neg_int:
  %v_data_ptr_int = getelementptr inbounds %Generic, %Generic* %v, i32 0, i32 1
  %v_data_int = load i8*, i8** %v_data_ptr_int
  %v_ptr_int = bitcast i8* %v_data_int to i64*
  %v_val_int = load i64, i64* %v_ptr_int
  %v_fval_int = sitofp i64 %v_val_int to double
  %v_neg_fval_int = fneg double %v_fval_int

  %v_val = fptosi double %v_neg_fval_int to i64
  
  ; Исправленное преобразование через временное хранение
  %v_val.i8.i_neg = inttoptr i64 %v_val to i8*
  %fresult.i_neg = call %Generic* @create(i32 0, i8* %v_val.i8.i_neg)
  ret %Generic* %fresult.i_neg
}

define %Generic* @not(%Generic* %v) {
entry:
  ; ИСПРАВЛЕННЫЙ ДОСТУП К ПОЛЯМ СТРУКТУРЫ
  %v_type_ptr = getelementptr inbounds %Generic, %Generic* %v, i32 0, i32 0
  %v_type = load i32, i32* %v_type_ptr
  
  ; Проверка типов
  %type_eq_bool = icmp eq i32 %v_type, 3
  br i1 %type_eq_bool, label %not_bool, label %error

not_bool:
  %v_data_ptr = getelementptr inbounds %Generic, %Generic* %v, i32 0, i32 1
  %v_data = load i8*, i8** %v_data_ptr
  %v_val = load i8, i8* %v_data

  %is_true = icmp eq i8 %v_val, 1
  br i1 %is_true, label %not_true, label %not_false

not_true:
  %result.false = call %Generic* @create(i32 3, i8* inttoptr (i8 0 to i8*))
  ret %Generic* %result.false

not_false:
  %result.true = call %Generic* @create(i32 3, i8* inttoptr (i8 1 to i8*))
  ret %Generic* %result.true

error:
  ret %Generic* null
}

define %Generic* @string_len(%Generic* %v) {
entry:
  ; ИСПРАВЛЕННЫЙ ДОСТУП К ПОЛЯМ СТРУКТУРЫ
  %v_type_ptr = getelementptr inbounds %Generic, %Generic* %v, i32 0, i32 0
  %v_type = load i32, i32* %v_type_ptr
  
  ; Проверка типов
  %type_eq_string = icmp eq i32 %v_type, 2
  br i1 %type_eq_string, label %len_string, label %error

len_string:
  %v_data_ptr = getelementptr inbounds %Generic, %Generic* %v, i32 0, i32 1
  %v_data = load i8*, i8** %v_data_ptr
  ; %v_val = bitcast i8* %v_data to i8*
  %v_len = call i64 @strlen(i8* %v_data)

  %len_ptr = inttoptr i64 %v_len to i8*
  %result = call %Generic* @create(i32 0, i8* %len_ptr)
  ret %Generic* %result

error:
  ret %Generic* null
}

define %Generic* @and(%Generic* %a, %Generic* %b) {
entry:
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  ; Проверка одинаковости типов
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %same_type, label %error

same_type:
  %type_bool = icmp eq i32 %a_type, 3
  br i1 %type_bool, label %and_bool, label %error

and_bool:
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  %a_val = load i8, i8* %a_data
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  %b_val = load i8, i8* %b_data
  
  %and = and i8 %a_val, %b_val
  %and.i8 = inttoptr i8 %and to i8*
  %result = call %Generic* @create(i32 3, i8* %and.i8)
  ret %Generic* %result
  
error:
  ret %Generic* null
}

define %Generic* @or(%Generic* %a, %Generic* %b) {
entry:
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  ; Проверка одинаковости типов
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %same_type, label %error

same_type:
  %type_bool = icmp eq i32 %a_type, 3
  br i1 %type_bool, label %or_bool, label %error

or_bool:
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  %a_val = load i8, i8* %a_data
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  %b_val = load i8, i8* %b_data
  
  %or = or i8 %a_val, %b_val
  %or.i8 = inttoptr i8 %or to i8*
  %result = call %Generic* @create(i32 3, i8* %or.i8)
  ret %Generic* %result
  
error:
  ret %Generic* null
}

define %Generic* @add(%Generic* %a, %Generic* %b) {
entry:
  ; ИСПРАВЛЕННЫЙ ДОСТУП К ПОЛЯМ СТРУКТУРЫ
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  ; Проверка одинаковости типов
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %same_type, label %add_mixed

add_mixed:
  switch i32 %a_type, label %error [
    i32 0, label %add_mixed_float_second
    i32 1, label %add_mixed_float_first
  ]

same_type:
  switch i32 %a_type, label %error [
    i32 0, label %add_int
    i32 1, label %add_float
  ]

add_int:
  ; ИСПРАВЛЕННЫЙ ДОСТУП ДЛЯ %b
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  %a_ptr = bitcast i8* %a_data to i64*
  %a_val = load i64, i64* %a_ptr
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  %b_ptr = bitcast i8* %b_data to i64*
  %b_val = load i64, i64* %b_ptr
  
  %sum = add i64 %a_val, %b_val
  %sum.i8 = inttoptr i64 %sum to i8*
  %result = call %Generic* @create(i32 0, i8* %sum.i8)
  ret %Generic* %result

add_float:
  %a_fdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata = load i8*, i8** %a_fdata_ptr
  %a_fptr = bitcast i8* %a_fdata to double*
  %a_fval = load double, double* %a_fptr
  
  %b_fdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata = load i8*, i8** %b_fdata_ptr
  %b_fptr = bitcast i8* %b_fdata to double*
  %b_fval = load double, double* %b_fptr
  
  %fsum = fadd double %a_fval, %b_fval
  
  ; Исправленное преобразование через временное хранение
  %temp.storage = alloca double
  store double %fsum, double* %temp.storage
  %as.i8 = bitcast double* %temp.storage to i8*
  
  %fresult = call %Generic* @create(i32 1, i8* %as.i8)
  ret %Generic* %fresult

add_mixed_float_first:
  %a_fdata_ptr.mixed_float_first = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata.mixed_float_first = load i8*, i8** %a_fdata_ptr.mixed_float_first
  %a_fptr.mixed_float_first = bitcast i8* %a_fdata.mixed_float_first to double*
  %a_fval.mixed_float_first = load double, double* %a_fptr.mixed_float_first
  
  %b_data_ptr.mixed_float_first = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data.mixed_float_first = load i8*, i8** %b_data_ptr.mixed_float_first
  %b_ptr.mixed_float_first = bitcast i8* %b_data.mixed_float_first to i64*
  %b_val.mixed_float_first = load i64, i64* %b_ptr.mixed_float_first
  %b_fval.mixed_float_first = sitofp i64 %b_val.mixed_float_first to double
  
  %fmixedsum.mixed_float_first = fadd double %a_fval.mixed_float_first, %b_fval.mixed_float_first
  
  ; Исправленное преобразование через временное хранение
  %temp.storage.mixed_float_first = alloca double
  store double %fmixedsum.mixed_float_first, double* %temp.storage.mixed_float_first
  %as.i8.mixed_float_first = bitcast double* %temp.storage.mixed_float_first to i8*
  
  %fresult.mixed_float_first = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_first)
  ret %Generic* %fresult.mixed_float_first

add_mixed_float_second:
  %a_data_ptr.mixed_float_second = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data.mixed_float_second = load i8*, i8** %a_data_ptr.mixed_float_second
  %a_ptr.mixed_float_second = bitcast i8* %a_data.mixed_float_second to i64*
  %a_val.mixed_float_second = load i64, i64* %a_ptr.mixed_float_second
  %a_fval.mixed_float_second = sitofp i64 %a_val.mixed_float_second to double 

  %b_fdata_ptr.mixed_float_second = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata.mixed_float_second = load i8*, i8** %b_fdata_ptr.mixed_float_second
  %b_fptr.mixed_float_second = bitcast i8* %b_fdata.mixed_float_second to double*
  %b_fval.mixed_float_second = load double, double* %b_fptr.mixed_float_second
  
  %fmixedsum.mixed_float_second = fadd double %a_fval.mixed_float_second, %b_fval.mixed_float_second
  
  ; Исправленное преобразование через временное хранение
  %temp.storage.mixed_float_second = alloca double
  store double %fmixedsum.mixed_float_second, double* %temp.storage.mixed_float_second
  %as.i8.mixed_float_second = bitcast double* %temp.storage.mixed_float_second to i8*
  
  %fresult.mixed_float_second = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_second)
  ret %Generic* %fresult.mixed_float_second
  
error:
  ret %Generic* null
}

define %Generic* @sub(%Generic* %a, %Generic* %b) {
entry:
  ; ИСПРАВЛЕННЫЙ ДОСТУП К ПОЛЯМ СТРУКТУРЫ
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  ; Проверка одинаковости типов
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %same_type, label %sub_mixed

same_type:
  switch i32 %a_type, label %error [
    i32 0, label %sub_int
    i32 1, label %sub_float
  ]

sub_mixed:
  switch i32 %a_type, label %error [
    i32 0, label %sub_mixed_float_second
    i32 1, label %sub_mixed_float_first
  ]

sub_int:
  ; ИСПРАВЛЕННЫЙ ДОСТУП ДЛЯ %b
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  %a_ptr = bitcast i8* %a_data to i64*
  %a_val = load i64, i64* %a_ptr
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  %b_ptr = bitcast i8* %b_data to i64*
  %b_val = load i64, i64* %b_ptr
  
  %sum = sub i64 %a_val, %b_val
  %sum.i8 = inttoptr i64 %sum to i8*
  %result = call %Generic* @create(i32 0, i8* %sum.i8)
  ret %Generic* %result

sub_float:
  %a_fdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata = load i8*, i8** %a_fdata_ptr
  %a_fptr = bitcast i8* %a_fdata to double*
  %a_fval = load double, double* %a_fptr
  
  %b_fdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata = load i8*, i8** %b_fdata_ptr
  %b_fptr = bitcast i8* %b_fdata to double*
  %b_fval = load double, double* %b_fptr
  
  %fsum = fsub double %a_fval, %b_fval
  
  ; Исправленное преобразование через временное хранение
  %temp.storage = alloca double
  store double %fsum, double* %temp.storage
  %as.i8 = bitcast double* %temp.storage to i8*
  
  %fresult = call %Generic* @create(i32 1, i8* %as.i8)
  ret %Generic* %fresult

sub_mixed_float_first:
  %a_fdata_ptr.mixed_float_first = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata.mixed_float_first = load i8*, i8** %a_fdata_ptr.mixed_float_first
  %a_fptr.mixed_float_first = bitcast i8* %a_fdata.mixed_float_first to double*
  %a_fval.mixed_float_first = load double, double* %a_fptr.mixed_float_first
  
  %b_data_ptr.mixed_float_first = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data.mixed_float_first = load i8*, i8** %b_data_ptr.mixed_float_first
  %b_ptr.mixed_float_first = bitcast i8* %b_data.mixed_float_first to i64*
  %b_val.mixed_float_first = load i64, i64* %b_ptr.mixed_float_first
  %b_fval.mixed_float_first = sitofp i64 %b_val.mixed_float_first to double
  
  %fmixedsum.mixed_float_first = fsub double %a_fval.mixed_float_first, %b_fval.mixed_float_first
  
  ; Исправленное преобразование через временное хранение
  %temp.storage.mixed_float_first = alloca double
  store double %fmixedsum.mixed_float_first, double* %temp.storage.mixed_float_first
  %as.i8.mixed_float_first = bitcast double* %temp.storage.mixed_float_first to i8*
  
  %fresult.mixed_float_first = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_first)
  ret %Generic* %fresult.mixed_float_first

sub_mixed_float_second:
  %a_data_ptr.mixed_float_second = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data.mixed_float_second = load i8*, i8** %a_data_ptr.mixed_float_second
  %a_ptr.mixed_float_second = bitcast i8* %a_data.mixed_float_second to i64*
  %a_val.mixed_float_second = load i64, i64* %a_ptr.mixed_float_second
  %a_fval.mixed_float_second = sitofp i64 %a_val.mixed_float_second to double 

  %b_fdata_ptr.mixed_float_second = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata.mixed_float_second = load i8*, i8** %b_fdata_ptr.mixed_float_second
  %b_fptr.mixed_float_second = bitcast i8* %b_fdata.mixed_float_second to double*
  %b_fval.mixed_float_second = load double, double* %b_fptr.mixed_float_second
  
  %fmixedsum.mixed_float_second = fsub double %a_fval.mixed_float_second, %b_fval.mixed_float_second
  
  ; Исправленное преобразование через временное хранение
  %temp.storage.mixed_float_second = alloca double
  store double %fmixedsum.mixed_float_second, double* %temp.storage.mixed_float_second
  %as.i8.mixed_float_second = bitcast double* %temp.storage.mixed_float_second to i8*
  
  %fresult.mixed_float_second = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_second)
  ret %Generic* %fresult.mixed_float_second
  
error:
  ret %Generic* null
}

define %Generic* @mul(%Generic* %a, %Generic* %b) {
entry:
  ; ИСПРАВЛЕННЫЙ ДОСТУП К ПОЛЯМ СТРУКТУРЫ
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  ; Проверка одинаковости типов
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %same_type, label %mul_mixed

mul_mixed:
  switch i32 %a_type, label %error [
    i32 0, label %mul_mixed_float_second
    i32 1, label %mul_mixed_float_first
  ]

same_type:
  switch i32 %a_type, label %error [
    i32 0, label %mul_int
    i32 1, label %mul_float
  ]

mul_int:
  ; ИСПРАВЛЕННЫЙ ДОСТУП ДЛЯ %b
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  %a_ptr = bitcast i8* %a_data to i64*
  %a_val = load i64, i64* %a_ptr
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  %b_ptr = bitcast i8* %b_data to i64*
  %b_val = load i64, i64* %b_ptr
  
  %sum = mul i64 %a_val, %b_val
  %sum.i8 = inttoptr i64 %sum to i8*
  %result = call %Generic* @create(i32 0, i8* %sum.i8)
  ret %Generic* %result

mul_float:
  %a_fdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata = load i8*, i8** %a_fdata_ptr
  %a_fptr = bitcast i8* %a_fdata to double*
  %a_fval = load double, double* %a_fptr
  
  %b_fdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata = load i8*, i8** %b_fdata_ptr
  %b_fptr = bitcast i8* %b_fdata to double*
  %b_fval = load double, double* %b_fptr
  
  %fsum = fmul double %a_fval, %b_fval
  
  ; Исправленное преобразование через временное хранение
  %temp.storage = alloca double
  store double %fsum, double* %temp.storage
  %as.i8 = bitcast double* %temp.storage to i8*
  
  %fresult = call %Generic* @create(i32 1, i8* %as.i8)
  ret %Generic* %fresult

mul_mixed_float_first:
  %a_fdata_ptr.mixed_float_first = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata.mixed_float_first = load i8*, i8** %a_fdata_ptr.mixed_float_first
  %a_fptr.mixed_float_first = bitcast i8* %a_fdata.mixed_float_first to double*
  %a_fval.mixed_float_first = load double, double* %a_fptr.mixed_float_first
  
  %b_data_ptr.mixed_float_first = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data.mixed_float_first = load i8*, i8** %b_data_ptr.mixed_float_first
  %b_ptr.mixed_float_first = bitcast i8* %b_data.mixed_float_first to i64*
  %b_val.mixed_float_first = load i64, i64* %b_ptr.mixed_float_first
  %b_fval.mixed_float_first = sitofp i64 %b_val.mixed_float_first to double
  
  %fmixedsum.mixed_float_first = fmul double %a_fval.mixed_float_first, %b_fval.mixed_float_first
  
  ; Исправленное преобразование через временное хранение
  %temp.storage.mixed_float_first = alloca double
  store double %fmixedsum.mixed_float_first, double* %temp.storage.mixed_float_first
  %as.i8.mixed_float_first = bitcast double* %temp.storage.mixed_float_first to i8*
  
  %fresult.mixed_float_first = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_first)
  ret %Generic* %fresult.mixed_float_first

mul_mixed_float_second:
  %a_data_ptr.mixed_float_second = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data.mixed_float_second = load i8*, i8** %a_data_ptr.mixed_float_second
  %a_ptr.mixed_float_second = bitcast i8* %a_data.mixed_float_second to i64*
  %a_val.mixed_float_second = load i64, i64* %a_ptr.mixed_float_second
  %a_fval.mixed_float_second = sitofp i64 %a_val.mixed_float_second to double 

  %b_fdata_ptr.mixed_float_second = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata.mixed_float_second = load i8*, i8** %b_fdata_ptr.mixed_float_second
  %b_fptr.mixed_float_second = bitcast i8* %b_fdata.mixed_float_second to double*
  %b_fval.mixed_float_second = load double, double* %b_fptr.mixed_float_second
  
  %fmixedsum.mixed_float_second = fmul double %a_fval.mixed_float_second, %b_fval.mixed_float_second
  
  ; Исправленное преобразование через временное хранение
  %temp.storage.mixed_float_second = alloca double
  store double %fmixedsum.mixed_float_second, double* %temp.storage.mixed_float_second
  %as.i8.mixed_float_second = bitcast double* %temp.storage.mixed_float_second to i8*
  
  %fresult.mixed_float_second = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_second)
  ret %Generic* %fresult.mixed_float_second
  
error:
  ret %Generic* null
}

define %Generic* @div(%Generic* %a, %Generic* %b) {
entry:
  ; ИСПРАВЛЕННЫЙ ДОСТУП К ПОЛЯМ СТРУКТУРЫ
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  ; Проверка одинаковости типов
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %same_type, label %div_mixed

div_mixed:
  switch i32 %a_type, label %error [
    i32 0, label %div_mixed_float_second
    i32 1, label %div_mixed_float_first
  ]

same_type:
  switch i32 %a_type, label %error [
    i32 0, label %div_int
    i32 1, label %div_float
  ]

div_int:
  ; ИСПРАВЛЕННЫЙ ДОСТУП ДЛЯ %b
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  %a_ptr = bitcast i8* %a_data to i64*
  %a_val = load i64, i64* %a_ptr
  %a_fval.conv = sitofp i64 %a_val to double
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  %b_ptr = bitcast i8* %b_data to i64*
  %b_val = load i64, i64* %b_ptr
  %b_fval.conv = sitofp i64 %b_val to double

  %zero_div_int = fcmp oeq double %b_fval.conv, 0.0
  br i1 %zero_div_int, label %error, label %continue_int

continue_int:
  %sum = fdiv double %a_fval.conv, %b_fval.conv
  %temp.storage.int = alloca double
  store double %sum, double* %temp.storage.int
  %as.i8.int = bitcast double* %temp.storage.int to i8*
  
  %fresult.int = call %Generic* @create(i32 1, i8* %as.i8.int)
  ret %Generic* %fresult.int

div_float:
  %a_fdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata = load i8*, i8** %a_fdata_ptr
  %a_fptr = bitcast i8* %a_fdata to double*
  %a_fval = load double, double* %a_fptr
  
  %b_fdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata = load i8*, i8** %b_fdata_ptr
  %b_fptr = bitcast i8* %b_fdata to double*
  %b_fval = load double, double* %b_fptr

  %zero_div_float = fcmp oeq double %b_fval, 0.0
  br i1 %zero_div_float, label %error, label %continue_float

continue_float:
  %fsum = fdiv double %a_fval, %b_fval
  ; Исправленное преобразование через временное хранение
  %temp.storage = alloca double
  store double %fsum, double* %temp.storage
  %as.i8 = bitcast double* %temp.storage to i8*
  
  %fresult = call %Generic* @create(i32 1, i8* %as.i8)
  ret %Generic* %fresult

div_mixed_float_first:
  %a_fdata_ptr.mixed_float_first = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata.mixed_float_first = load i8*, i8** %a_fdata_ptr.mixed_float_first
  %a_fptr.mixed_float_first = bitcast i8* %a_fdata.mixed_float_first to double*
  %a_fval.mixed_float_first = load double, double* %a_fptr.mixed_float_first
  
  %b_data_ptr.mixed_float_first = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data.mixed_float_first = load i8*, i8** %b_data_ptr.mixed_float_first
  %b_ptr.mixed_float_first = bitcast i8* %b_data.mixed_float_first to i64*
  %b_val.mixed_float_first = load i64, i64* %b_ptr.mixed_float_first
  %b_fval.mixed_float_first = sitofp i64 %b_val.mixed_float_first to double

  %zero_div.mixed_float_first = fcmp oeq double %b_fval.mixed_float_first, 0.0
  br i1 %zero_div.mixed_float_first, label %error, label %continue_mixed_float_first

continue_mixed_float_first:
  %fmixedsum.mixed_float_first = fdiv double %a_fval.mixed_float_first, %b_fval.mixed_float_first

  ; Исправленное преобразование через временное хранение
  %temp.storage.mixed_float_first = alloca double
  store double %fmixedsum.mixed_float_first, double* %temp.storage.mixed_float_first
  %as.i8.mixed_float_first = bitcast double* %temp.storage.mixed_float_first to i8*
  
  %fresult.mixed_float_first = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_first)
  ret %Generic* %fresult.mixed_float_first

div_mixed_float_second:
  %a_data_ptr.mixed_float_second = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data.mixed_float_second = load i8*, i8** %a_data_ptr.mixed_float_second
  %a_ptr.mixed_float_second = bitcast i8* %a_data.mixed_float_second to i64*
  %a_val.mixed_float_second = load i64, i64* %a_ptr.mixed_float_second
  %a_fval.mixed_float_second = sitofp i64 %a_val.mixed_float_second to double 

  %b_fdata_ptr.mixed_float_second = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata.mixed_float_second = load i8*, i8** %b_fdata_ptr.mixed_float_second
  %b_fptr.mixed_float_second = bitcast i8* %b_fdata.mixed_float_second to double*
  %b_fval.mixed_float_second = load double, double* %b_fptr.mixed_float_second

  %zero_div.mixed_float_second = fcmp oeq double %b_fval.mixed_float_second, 0.0
  br i1 %zero_div.mixed_float_second, label %error, label %continue_mixed_float_second

continue_mixed_float_second:
  %fmixedsum.mixed_float_second = fdiv double %a_fval.mixed_float_second, %b_fval.mixed_float_second
  
  ; Исправленное преобразование через временное хранение
  %temp.storage.mixed_float_second = alloca double
  store double %fmixedsum.mixed_float_second, double* %temp.storage.mixed_float_second
  %as.i8.mixed_float_second = bitcast double* %temp.storage.mixed_float_second to i8*
  
  %fresult.mixed_float_second = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_second)
  ret %Generic* %fresult.mixed_float_second
  
error:
  ret %Generic* null
}

define %Generic* @mod(%Generic* %a, %Generic* %b) {
entry:
  ; ИСПРАВЛЕННЫЙ ДОСТУП К ПОЛЯМ СТРУКТУРЫ
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  ; Проверка одинаковости типов
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %same_type, label %error

same_type:
  switch i32 %a_type, label %error [
    i32 0, label %mod_int
    i32 1, label %error
  ]

mod_int:
  ; ИСПРАВЛЕННЫЙ ДОСТУП ДЛЯ %b
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  %a_ptr = bitcast i8* %a_data to i64*
  %a_val = load i64, i64* %a_ptr
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  %b_ptr = bitcast i8* %b_data to i64*
  %b_val = load i64, i64* %b_ptr

  %zero_div_int = icmp eq i64 %b_val, 0
  br i1 %zero_div_int, label %error, label %continue_int

continue_int:
  %sum = sdiv i64 %a_val, %b_val
  %sum.i8 = inttoptr i64 %sum to i8*
  %result = call %Generic* @create(i32 0, i8* %sum.i8)
  ret %Generic* %result
  
error:
  ret %Generic* null
}

define %Generic* @rem(%Generic* %a, %Generic* %b) {
entry:
  ; ИСПРАВЛЕННЫЙ ДОСТУП К ПОЛЯМ СТРУКТУРЫ
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  ; Проверка одинаковости типов
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %same_type, label %error

same_type:
  switch i32 %a_type, label %error [
    i32 0, label %rem_int
    i32 1, label %error
  ]

rem_int:
  ; ИСПРАВЛЕННЫЙ ДОСТУП ДЛЯ %b
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  %a_ptr = bitcast i8* %a_data to i64*
  %a_val = load i64, i64* %a_ptr
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  %b_ptr = bitcast i8* %b_data to i64*
  %b_val = load i64, i64* %b_ptr

  %zero_div_int = icmp eq i64 %b_val, 0
  br i1 %zero_div_int, label %error, label %continue_int

continue_int:
  %sum = srem i64 %a_val, %b_val
  %sum.i8 = inttoptr i64 %sum to i8*
  %result = call %Generic* @create(i32 0, i8* %sum.i8)
  ret %Generic* %result
  
error:
  ret %Generic* null
}

define %Generic* @power(%Generic* %a, %Generic* %b) {
entry:
  ; ИСПРАВЛЕННЫЙ ДОСТУП К ПОЛЯМ СТРУКТУРЫ
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  ; Проверка одинаковости типов
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %same_type, label %pow_mixed

pow_mixed:
  switch i32 %a_type, label %error [
    i32 0, label %pow_mixed_float_second
    i32 1, label %pow_mixed_float_first
  ]

same_type:
  switch i32 %a_type, label %error [
    i32 0, label %pow_int
    i32 1, label %pow_float
  ]

pow_int:
  ; ИСПРАВЛЕННЫЙ ДОСТУП ДЛЯ %b
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  %a_ptr = bitcast i8* %a_data to i64*
  %a_val = load i64, i64* %a_ptr
  %a_fval.conv = sitofp i64 %a_val to double
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  %b_ptr = bitcast i8* %b_data to i64*
  %b_val = load i64, i64* %b_ptr
  %b_fval.conv = sitofp i64 %b_val to double

  %sum = call double @pow(double %a_fval.conv, double %b_fval.conv)
  %temp.storage.int = alloca double
  store double %sum, double* %temp.storage.int
  %as.i8.int = bitcast double* %temp.storage.int to i8*
  
  %fresult.int = call %Generic* @create(i32 1, i8* %as.i8.int)
  ret %Generic* %fresult.int

pow_float:
  %a_fdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata = load i8*, i8** %a_fdata_ptr
  %a_fptr = bitcast i8* %a_fdata to double*
  %a_fval = load double, double* %a_fptr
  
  %b_fdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata = load i8*, i8** %b_fdata_ptr
  %b_fptr = bitcast i8* %b_fdata to double*
  %b_fval = load double, double* %b_fptr

  %fsum = call double @pow(double %a_fval, double %b_fval)
  ; Исправленное преобразование через временное хранение
  %temp.storage = alloca double
  store double %fsum, double* %temp.storage
  %as.i8 = bitcast double* %temp.storage to i8*
  
  %fresult = call %Generic* @create(i32 1, i8* %as.i8)
  ret %Generic* %fresult

pow_mixed_float_first:
  %a_fdata_ptr.mixed_float_first = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata.mixed_float_first = load i8*, i8** %a_fdata_ptr.mixed_float_first
  %a_fptr.mixed_float_first = bitcast i8* %a_fdata.mixed_float_first to double*
  %a_fval.mixed_float_first = load double, double* %a_fptr.mixed_float_first
  
  %b_data_ptr.mixed_float_first = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data.mixed_float_first = load i8*, i8** %b_data_ptr.mixed_float_first
  %b_ptr.mixed_float_first = bitcast i8* %b_data.mixed_float_first to i64*
  %b_val.mixed_float_first = load i64, i64* %b_ptr.mixed_float_first
  %b_fval.mixed_float_first = sitofp i64 %b_val.mixed_float_first to double

  %fmixedsum.mixed_float_first = call double @pow(double %a_fval.mixed_float_first, double %b_fval.mixed_float_first)

  ; Исправленное преобразование через временное хранение
  %temp.storage.mixed_float_first = alloca double
  store double %fmixedsum.mixed_float_first, double* %temp.storage.mixed_float_first
  %as.i8.mixed_float_first = bitcast double* %temp.storage.mixed_float_first to i8*
  
  %fresult.mixed_float_first = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_first)
  ret %Generic* %fresult.mixed_float_first

pow_mixed_float_second:
  %a_data_ptr.mixed_float_second = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data.mixed_float_second = load i8*, i8** %a_data_ptr.mixed_float_second
  %a_ptr.mixed_float_second = bitcast i8* %a_data.mixed_float_second to i64*
  %a_val.mixed_float_second = load i64, i64* %a_ptr.mixed_float_second
  %a_fval.mixed_float_second = sitofp i64 %a_val.mixed_float_second to double 

  %b_fdata_ptr.mixed_float_second = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata.mixed_float_second = load i8*, i8** %b_fdata_ptr.mixed_float_second
  %b_fptr.mixed_float_second = bitcast i8* %b_fdata.mixed_float_second to double*
  %b_fval.mixed_float_second = load double, double* %b_fptr.mixed_float_second

  %fmixedsum.mixed_float_second = call double @pow(double %a_fval.mixed_float_second, double %b_fval.mixed_float_second)
  
  ; Исправленное преобразование через временное хранение
  %temp.storage.mixed_float_second = alloca double
  store double %fmixedsum.mixed_float_second, double* %temp.storage.mixed_float_second
  %as.i8.mixed_float_second = bitcast double* %temp.storage.mixed_float_second to i8*
  
  %fresult.mixed_float_second = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_second)
  ret %Generic* %fresult.mixed_float_second
  
error:
  ret %Generic* null
}

define %Generic* @equal(%Generic* %a, %Generic* %b) {
entry:
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %check_value, label %error

check_value:
  switch i32 %a_type, label %error [
    i32 0, label %cmp_int
    i32 1, label %cmp_float
    i32 2, label %cmp_str
    i32 3, label %cmp_bool
  ]

cmp_int:
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  
  %a_ptr = bitcast i8* %a_data to i64*
  %a_val = load i64, i64* %a_ptr
  %b_ptr = bitcast i8* %b_data to i64*
  %b_val = load i64, i64* %b_ptr
  
  %cmp = icmp eq i64 %a_val, %b_val
  %cmp.int = inttoptr i1 %cmp to i8*
  %result.int = call %Generic* @create(i32 3, i8* %cmp.int)
  ret %Generic* %result.int

cmp_float:
  %a_fdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata = load i8*, i8** %a_fdata_ptr
  %a_fptr = bitcast i8* %a_fdata to double*
  %a_fval = load double, double* %a_fptr
  
  %b_fdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata = load i8*, i8** %b_fdata_ptr
  %b_fptr = bitcast i8* %b_fdata to double*
  %b_fval = load double, double* %b_fptr
  
  %fcmp = fcmp oeq double %a_fval, %b_fval
  %cmp.float = inttoptr i1 %fcmp to i8*
  %result.float = call %Generic* @create(i32 3, i8* %cmp.float)
  ret %Generic* %result.float

cmp_str:
  %a_sdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_str = load i8*, i8** %a_sdata_ptr
  %b_sdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_str = load i8*, i8** %b_sdata_ptr
  
  %cmp_res = call i32 @strcmp(i8* %a_str, i8* %b_str)
  %res.string = icmp eq i32 %cmp_res, 0
  %cmp.string = inttoptr i1 %res.string to i8*
  %result.string = call %Generic* @create(i32 3, i8* %cmp.string)
  ret %Generic* %result.string

cmp_bool:
  %a_bdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_bdata = load i8*, i8** %a_bdata_ptr
  %a_bval = load i8, i8* %a_bdata
  
  %b_bdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_bdata = load i8*, i8** %b_bdata_ptr
  %b_bval = load i8, i8* %b_bdata
  
  %cmp_res.bool = icmp eq i8 %a_bval, %b_bval
  %cmp.bool = inttoptr i1 %cmp_res.bool to i8*
  %result.bool = call %Generic* @create(i32 3, i8* %cmp.bool)
  ret %Generic* %result.bool
  
error:
  ret %Generic* null
}

define %Generic* @nequal(%Generic* %a, %Generic* %b) {
entry:
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %check_value, label %error

check_value:
  switch i32 %a_type, label %error [
    i32 0, label %cmp_int
    i32 1, label %cmp_float
    i32 2, label %cmp_str
    i32 3, label %cmp_bool
  ]

cmp_int:
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  
  %a_ptr = bitcast i8* %a_data to i64*
  %a_val = load i64, i64* %a_ptr
  %b_ptr = bitcast i8* %b_data to i64*
  %b_val = load i64, i64* %b_ptr
  
  %cmp = icmp ne i64 %a_val, %b_val
  %cmp.int = inttoptr i1 %cmp to i8*
  %result.int = call %Generic* @create(i32 3, i8* %cmp.int)
  ret %Generic* %result.int

cmp_float:
  %a_fdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata = load i8*, i8** %a_fdata_ptr
  %a_fptr = bitcast i8* %a_fdata to double*
  %a_fval = load double, double* %a_fptr
  
  %b_fdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata = load i8*, i8** %b_fdata_ptr
  %b_fptr = bitcast i8* %b_fdata to double*
  %b_fval = load double, double* %b_fptr
  
  %fcmp = fcmp one double %a_fval, %b_fval
  %cmp.float = inttoptr i1 %fcmp to i8*
  %result.float = call %Generic* @create(i32 3, i8* %cmp.float)
  ret %Generic* %result.float

cmp_str:
  %a_sdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_str = load i8*, i8** %a_sdata_ptr
  %b_sdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_str = load i8*, i8** %b_sdata_ptr
  
  %cmp_res = call i32 @strcmp(i8* %a_str, i8* %b_str)
  %res.string = icmp ne i32 %cmp_res, 0
  %cmp.string = inttoptr i1 %res.string to i8*
  %result.string = call %Generic* @create(i32 3, i8* %cmp.string)
  ret %Generic* %result.string

cmp_bool:
  %a_bdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_bdata = load i8*, i8** %a_bdata_ptr
  %a_bval = load i8, i8* %a_bdata
  
  %b_bdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_bdata = load i8*, i8** %b_bdata_ptr
  %b_bval = load i8, i8* %b_bdata
  
  %cmp_res.bool = icmp ne i8 %a_bval, %b_bval
  %cmp.bool = inttoptr i1 %cmp_res.bool to i8*
  %result.bool = call %Generic* @create(i32 3, i8* %cmp.bool)
  ret %Generic* %result.bool
  
error:
  ret %Generic* null
}

define %Generic* @gt(%Generic* %a, %Generic* %b) {
entry:
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %check_value, label %error

check_value:
  switch i32 %a_type, label %error [
    i32 0, label %cmp_int
    i32 1, label %cmp_float
    i32 2, label %cmp_str
    i32 3, label %cmp_bool
  ]

cmp_int:
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  
  %a_ptr = bitcast i8* %a_data to i64*
  %a_val = load i64, i64* %a_ptr
  %b_ptr = bitcast i8* %b_data to i64*
  %b_val = load i64, i64* %b_ptr
  
  %cmp = icmp sgt i64 %a_val, %b_val
  %cmp.int = inttoptr i1 %cmp to i8*
  %result.int = call %Generic* @create(i32 3, i8* %cmp.int)
  ret %Generic* %result.int

cmp_float:
  %a_fdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata = load i8*, i8** %a_fdata_ptr
  %a_fptr = bitcast i8* %a_fdata to double*
  %a_fval = load double, double* %a_fptr
  
  %b_fdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata = load i8*, i8** %b_fdata_ptr
  %b_fptr = bitcast i8* %b_fdata to double*
  %b_fval = load double, double* %b_fptr
  
  %fcmp = fcmp ogt double %a_fval, %b_fval
  %cmp.float = inttoptr i1 %fcmp to i8*
  %result.float = call %Generic* @create(i32 3, i8* %cmp.float)
  ret %Generic* %result.float

cmp_str:
  %a_sdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_str = load i8*, i8** %a_sdata_ptr
  %b_sdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_str = load i8*, i8** %b_sdata_ptr
  
  %cmp_res = call i32 @strcmp(i8* %a_str, i8* %b_str)
  %res.string = icmp sgt i32 %cmp_res, 0
  %cmp.string = inttoptr i1 %res.string to i8*
  %result.string = call %Generic* @create(i32 3, i8* %cmp.string)
  ret %Generic* %result.string

cmp_bool:
  %a_bdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_bdata = load i8*, i8** %a_bdata_ptr
  %a_bval = load i8, i8* %a_bdata
  
  %b_bdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_bdata = load i8*, i8** %b_bdata_ptr
  %b_bval = load i8, i8* %b_bdata
  
  %cmp_res.bool = icmp sgt i8 %a_bval, %b_bval
  %cmp.bool = inttoptr i1 %cmp_res.bool to i8*
  %result.bool = call %Generic* @create(i32 3, i8* %cmp.bool)
  ret %Generic* %result.bool
  
error:
  ret %Generic* null
}

define %Generic* @ge(%Generic* %a, %Generic* %b) {
entry:
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %check_value, label %error

check_value:
  switch i32 %a_type, label %error [
    i32 0, label %cmp_int
    i32 1, label %cmp_float
    i32 2, label %cmp_str
    i32 3, label %cmp_bool
  ]

cmp_int:
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  
  %a_ptr = bitcast i8* %a_data to i64*
  %a_val = load i64, i64* %a_ptr
  %b_ptr = bitcast i8* %b_data to i64*
  %b_val = load i64, i64* %b_ptr
  
  %cmp = icmp sge i64 %a_val, %b_val
  %cmp.int = inttoptr i1 %cmp to i8*
  %result.int = call %Generic* @create(i32 3, i8* %cmp.int)
  ret %Generic* %result.int

cmp_float:
  %a_fdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata = load i8*, i8** %a_fdata_ptr
  %a_fptr = bitcast i8* %a_fdata to double*
  %a_fval = load double, double* %a_fptr
  
  %b_fdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata = load i8*, i8** %b_fdata_ptr
  %b_fptr = bitcast i8* %b_fdata to double*
  %b_fval = load double, double* %b_fptr
  
  %fcmp = fcmp oge double %a_fval, %b_fval
  %cmp.float = inttoptr i1 %fcmp to i8*
  %result.float = call %Generic* @create(i32 3, i8* %cmp.float)
  ret %Generic* %result.float

cmp_str:
  %a_sdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_str = load i8*, i8** %a_sdata_ptr
  %b_sdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_str = load i8*, i8** %b_sdata_ptr
  
  %cmp_res = call i32 @strcmp(i8* %a_str, i8* %b_str)
  %res.string.less = icmp sgt i32 %cmp_res, 0
  %res.string.equal = icmp eq i32 %cmp_res, 0

  switch i1 %res.string.less, label %error [
    i1 0, label %equal_string
    i1 1, label %geq_string
  ]

equal_string:
  switch i1 %res.string.equal, label %error [
    i1 0, label %ngeq_string
    i1 1, label %geq_string
  ]

geq_string:
  %cmp.string.geq = inttoptr i1 1 to i8*
  %result.string.geq = call %Generic* @create(i32 3, i8* %cmp.string.geq)
  ret %Generic* %result.string.geq

ngeq_string:
  %cmp.string.ngeq = inttoptr i1 0 to i8*
  %result.string.ngeq = call %Generic* @create(i32 3, i8* %cmp.string.ngeq)
  ret %Generic* %result.string.ngeq

cmp_bool:
  %a_bdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_bdata = load i8*, i8** %a_bdata_ptr
  %a_bval = load i8, i8* %a_bdata
  
  %b_bdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_bdata = load i8*, i8** %b_bdata_ptr
  %b_bval = load i8, i8* %b_bdata
  
  %cmp_res.bool = icmp sge i8 %a_bval, %b_bval
  %cmp.bool = inttoptr i1 %cmp_res.bool to i8*
  %result.bool = call %Generic* @create(i32 3, i8* %cmp.bool)
  ret %Generic* %result.bool
  
error:
  ret %Generic* null
}

define %Generic* @lt(%Generic* %a, %Generic* %b) {
entry:
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %check_value, label %error

check_value:
  switch i32 %a_type, label %error [
    i32 0, label %cmp_int
    i32 1, label %cmp_float
    i32 2, label %cmp_str
    i32 3, label %cmp_bool
  ]

cmp_int:
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  
  %a_ptr = bitcast i8* %a_data to i64*
  %a_val = load i64, i64* %a_ptr
  %b_ptr = bitcast i8* %b_data to i64*
  %b_val = load i64, i64* %b_ptr
  
  %cmp = icmp slt i64 %a_val, %b_val
  %cmp.int = inttoptr i1 %cmp to i8*
  %result.int = call %Generic* @create(i32 3, i8* %cmp.int)
  ret %Generic* %result.int

cmp_float:
  %a_fdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata = load i8*, i8** %a_fdata_ptr
  %a_fptr = bitcast i8* %a_fdata to double*
  %a_fval = load double, double* %a_fptr
  
  %b_fdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata = load i8*, i8** %b_fdata_ptr
  %b_fptr = bitcast i8* %b_fdata to double*
  %b_fval = load double, double* %b_fptr
  
  %fcmp = fcmp olt double %a_fval, %b_fval
  %cmp.float = inttoptr i1 %fcmp to i8*
  %result.float = call %Generic* @create(i32 3, i8* %cmp.float)
  ret %Generic* %result.float

cmp_str:
  %a_sdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_str = load i8*, i8** %a_sdata_ptr
  %b_sdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_str = load i8*, i8** %b_sdata_ptr
  
  %cmp_res = call i32 @strcmp(i8* %a_str, i8* %b_str)
  %res.string = icmp slt i32 %cmp_res, 0
  %cmp.string = inttoptr i1 %res.string to i8*
  %result.string = call %Generic* @create(i32 3, i8* %cmp.string)
  ret %Generic* %result.string

cmp_bool:
  %a_bdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_bdata = load i8*, i8** %a_bdata_ptr
  %a_bval = load i8, i8* %a_bdata
  
  %b_bdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_bdata = load i8*, i8** %b_bdata_ptr
  %b_bval = load i8, i8* %b_bdata
  
  %cmp_res.bool = icmp slt i8 %a_bval, %b_bval
  %cmp.bool = inttoptr i1 %cmp_res.bool to i8*
  %result.bool = call %Generic* @create(i32 3, i8* %cmp.bool)
  ret %Generic* %result.bool
  
error:
  ret %Generic* null
}

define %Generic* @le(%Generic* %a, %Generic* %b) {
entry:
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %check_value, label %error

check_value:
  switch i32 %a_type, label %error [
    i32 0, label %cmp_int
    i32 1, label %cmp_float
    i32 2, label %cmp_str
    i32 3, label %cmp_bool
  ]

cmp_int:
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_data = load i8*, i8** %a_data_ptr
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_data = load i8*, i8** %b_data_ptr
  
  %a_ptr = bitcast i8* %a_data to i64*
  %a_val = load i64, i64* %a_ptr
  %b_ptr = bitcast i8* %b_data to i64*
  %b_val = load i64, i64* %b_ptr
  
  %cmp = icmp sle i64 %a_val, %b_val
  %cmp.int = inttoptr i1 %cmp to i8*
  %result.int = call %Generic* @create(i32 3, i8* %cmp.int)
  ret %Generic* %result.int

cmp_float:
  %a_fdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_fdata = load i8*, i8** %a_fdata_ptr
  %a_fptr = bitcast i8* %a_fdata to double*
  %a_fval = load double, double* %a_fptr
  
  %b_fdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_fdata = load i8*, i8** %b_fdata_ptr
  %b_fptr = bitcast i8* %b_fdata to double*
  %b_fval = load double, double* %b_fptr
  
  %fcmp = fcmp ole double %a_fval, %b_fval
  %cmp.float = inttoptr i1 %fcmp to i8*
  %result.float = call %Generic* @create(i32 3, i8* %cmp.float)
  ret %Generic* %result.float

cmp_str:
  %a_sdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_str = load i8*, i8** %a_sdata_ptr
  %b_sdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_str = load i8*, i8** %b_sdata_ptr
  
  %cmp_res = call i32 @strcmp(i8* %a_str, i8* %b_str)
  %res.string.less = icmp slt i32 %cmp_res, 0
  %res.string.equal = icmp eq i32 %cmp_res, 0

  switch i1 %res.string.less, label %error [
    i1 0, label %equal_string
    i1 1, label %leq_string
  ]

equal_string:
  switch i1 %res.string.equal, label %error [
    i1 0, label %nleq_string
    i1 1, label %leq_string
  ]

leq_string:
  %cmp.string.leq = inttoptr i1 1 to i8*
  %result.string.leq = call %Generic* @create(i32 3, i8* %cmp.string.leq)
  ret %Generic* %result.string.leq

nleq_string:
  %cmp.string.nleq = inttoptr i1 0 to i8*
  %result.string.nleq = call %Generic* @create(i32 3, i8* %cmp.string.nleq)
  ret %Generic* %result.string.nleq

cmp_bool:
  %a_bdata_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_bdata = load i8*, i8** %a_bdata_ptr
  %a_bval = load i8, i8* %a_bdata
  
  %b_bdata_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_bdata = load i8*, i8** %b_bdata_ptr
  %b_bval = load i8, i8* %b_bdata
  
  %cmp_res.bool = icmp sle i8 %a_bval, %b_bval
  %cmp.bool = inttoptr i1 %cmp_res.bool to i8*
  %result.bool = call %Generic* @create(i32 3, i8* %cmp.bool)
  ret %Generic* %result.bool
  
error:
  ret %Generic* null
}

define void @destroy(%Generic* %obj) {
entry:
  ; ИСПРАВЛЕННЫЙ ДОСТУП К ТИПУ
  %type_ptr = getelementptr inbounds %Generic, %Generic* %obj, i32 0, i32 0
  %type = load i32, i32* %type_ptr
  
  ; ИСПРАВЛЕННЫЙ ДОСТУП К ДАННЫМ
  %data_ptr_ptr = getelementptr inbounds %Generic, %Generic* %obj, i32 0, i32 1
  %data_ptr = load i8*, i8** %data_ptr_ptr
  
  switch i32 %type, label %free_obj [
    i32 2, label %free_str
  ]

free_str:
  call void @free(i8* %data_ptr)
  br label %free_obj

free_obj:
  %obj_ptr = bitcast %Generic* %obj to i8*
  call void @free(i8* %obj_ptr)
  ret void
}

; Функция печати
define void @print(%Generic* %obj) {
entry:
  %is_null = icmp eq %Generic* %obj, null
  br i1 %is_null, label %print_nil, label %print_value

print_nil:
  %nil_fmt = getelementptr inbounds [5 x i8], [5 x i8]* @.str.nil, i32 0, i32 0
  call i32 (i8*, ...) @printf(i8* %nil_fmt)
  ret void

print_value:
  %type_ptr = getelementptr inbounds %Generic, %Generic* %obj, i32 0, i32 0
  %type = load i32, i32* %type_ptr
  
  %data_ptr_ptr = getelementptr inbounds %Generic, %Generic* %obj, i32 0, i32 1
  %data_ptr = load i8*, i8** %data_ptr_ptr
  
  switch i32 %type, label %unknown [
    i32 0, label %print_int
    i32 1, label %print_float
    i32 2, label %print_string
    i32 3, label %print_bool
  ]

print_int:
  %int_ptr = bitcast i8* %data_ptr to i64*
  %int_val = load i64, i64* %int_ptr
  %int_fmt = getelementptr inbounds [6 x i8], [6 x i8]* @.str.int, i32 0, i32 0
  call i32 (i8*, ...) @printf(i8* %int_fmt, i64 %int_val)
  ret void

print_float:
  %float_ptr = bitcast i8* %data_ptr to double*
  %float_val = load double, double* %float_ptr
  %float_fmt = getelementptr inbounds [4 x i8], [4 x i8]* @.str.float, i32 0, i32 0
  call i32 (i8*, ...) @printf(i8* %float_fmt, double %float_val)
  ret void

print_string:
  %string_val = bitcast i8* %data_ptr to i8*
  %string_fmt = getelementptr inbounds [4 x i8], [4 x i8]* @.str.string, i32 0, i32 0
  call i32 (i8*, ...) @printf(i8* %string_fmt, i8* %string_val)
  ret void

print_bool:
  %bool_ptr = bitcast i8* %data_ptr to i8*
  %bool_val = load i8, i8* %bool_ptr

  %is_true = icmp eq i32 %bool_val, 1
  br i1 %is_true, label %print_true, label %print_false

print_true:
  %true_fmt = getelementptr inbounds [6 x i8], [6 x i8]* @.str.true, i32 0, i32 0
  call i32 (i8*, ...) @printf(i8* %true_fmt)
  ret void

print_false:
  %false_fmt = getelementptr inbounds [7 x i8], [7 x i8]* @.str.false, i32 0, i32 0
  call i32 (i8*, ...) @printf(i8* %false_fmt)
  ret void

unknown:
  ret void
}

define %Generic* @concat(%Generic* %a, %Generic* %b) {
entry:
  ; Получаем тип первого аргумента
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  ; Получаем тип второго аргумента
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  ; Проверяем что оба типа - строки (тип 2)
  %both_string = icmp eq i32 %a_type, 2
  %both_string1 = icmp eq i32 %b_type, 2
  %both_ok = and i1 %both_string, %both_string1
  br i1 %both_ok, label %concat_strings, label %error

concat_strings:
  ; Получаем указатели на строки
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_str_ptr = load i8*, i8** %a_data_ptr
  %a_str = load i8*, i8** %a_data_ptr
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_str = load i8*, i8** %b_data_ptr
  
  ; Вычисляем длину результирующей строки
  %a_len = call i64 @strlen(i8* %a_str)
  %b_len = call i64 @strlen(i8* %b_str)
  %total_len = add i64 %a_len, %b_len
  %buffer = call i8* @malloc(i64 %total_len)
  
  ; Копируем данные
  call i8* @strcpy(i8* %buffer, i8* %a_str)
  call i8* @strcat(i8* %buffer, i8* %b_str)
  
  ; Создаем новый объект
  %result = call %Generic* @create(i32 2, i8* %buffer)
  call void @free(i8* %buffer)
  ret %Generic* %result

error:
  ret %Generic* null
}

