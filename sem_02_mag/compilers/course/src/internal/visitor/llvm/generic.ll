%Generic = type {
  i32,     ; тип данных (0=int, 1=float, 2=string, 3=bool, 4=table, 5=nil)
  i8*      ; указатель на данные
}

@NIL_TYPE = constant i32 5

@.error.null_value = private constant [11 x i8] c"Null value\00"

define %Generic* @create_nil() {
entry:
  %size = ptrtoint %Generic* getelementptr inbounds (%Generic, %Generic* null, i32 1) to i64
  
  %nil = call i8* @malloc(i64 %size)
  %nil_generic = bitcast i8* %nil to %Generic*
  
  %type_ptr = getelementptr inbounds %Generic, %Generic* %nil_generic, i32 0, i32 0
  store i32 5, i32* %type_ptr, align 4
  
  %data_ptr = getelementptr inbounds %Generic, %Generic* %nil_generic, i32 0, i32 1
  store i8* null, i8** %data_ptr, align 8
  
  ret %Generic* %nil_generic
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
    i32 4, label %init_table
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
  %len_str = call i64 @strlen(i8* %value)
  %len = add i64 %len_str, 1
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

init_table:
  %data_table = getelementptr inbounds %Generic, %Generic* %g, i32 0, i32 1
  store i8* %value, i8** %data_table
  ret %Generic* %g

error:
  call void @panic(i8* getelementptr inbounds ([11 x i8], [11 x i8]* @.error.null_value, i32 0, i32 0))
  ret %Generic* %g

invalid:
  ret %Generic* %g
}

define void @copy(%Generic* %src, %Generic* %dst) {
entry:
  %src.type.ptr = getelementptr %Generic, %Generic* %src, i32 0, i32 0
  %type = load i32, i32* %src.type.ptr
  %dst.type.ptr = getelementptr %Generic, %Generic* %dst, i32 0, i32 0
  store i32 %type, i32* %dst.type.ptr

  %src.data.ptr = getelementptr %Generic, %Generic* %src, i32 0, i32 1
  %data = load i8*, i8** %src.data.ptr
  %dst.data.ptr = getelementptr %Generic, %Generic* %dst, i32 0, i32 1
  store i8* %data, i8** %dst.data.ptr

  ret void
}

define void @destroy(%Generic* %obj) {
entry:
  %type_ptr = getelementptr inbounds %Generic, %Generic* %obj, i32 0, i32 0
  %type = load i32, i32* %type_ptr
  
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