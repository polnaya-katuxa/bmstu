@.error.string_expected = private constant [23 x i8] c"Value should be string\00"

define %Generic* @string_len(%Generic* %v) {
entry:
  %v_type_ptr = getelementptr inbounds %Generic, %Generic* %v, i32 0, i32 0
  %v_type = load i32, i32* %v_type_ptr
  
  %type_eq_string = icmp eq i32 %v_type, 2
  br i1 %type_eq_string, label %len_string, label %error

len_string:
  %v_data_ptr = getelementptr inbounds %Generic, %Generic* %v, i32 0, i32 1
  %v_data = load i8*, i8** %v_data_ptr
  %v_len = call i64 @strlen(i8* %v_data)

  %len_ptr = inttoptr i64 %v_len to i8*
  %result = call %Generic* @create(i32 0, i8* %len_ptr)
  ret %Generic* %result

error:
  call void @panic(i8* getelementptr inbounds ([23 x i8], [23 x i8]* @.error.string_expected, i32 0, i32 0))
  ret %Generic* null
}

declare void @llvm.memset.p0i8.i64(i8* nocapture writeonly, i8, i64, i1 immarg)

define %Generic* @concat(%Generic* %a, %Generic* %b) {
entry:
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  %both_string = icmp eq i32 %a_type, 2
  %both_string1 = icmp eq i32 %b_type, 2
  %both_ok = and i1 %both_string, %both_string1
  br i1 %both_ok, label %concat_strings, label %error

concat_strings:
  %a_data_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
  %a_str_ptr = load i8*, i8** %a_data_ptr
  %a_str = load i8*, i8** %a_data_ptr
  
  %b_data_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
  %b_str = load i8*, i8** %b_data_ptr
  
  %a_len = call i64 @strlen(i8* %a_str)
  %b_len = call i64 @strlen(i8* %b_str)
  %sum_len = add i64 %a_len, %b_len
  %total_len = add i64 %sum_len, 1
  %buffer = call i8* @malloc(i64 %total_len)

  call void @llvm.memset.p0i8.i64(i8* %buffer, i8 0, i64 %total_len, i1 false)
  
  call i8* @strcpy(i8* %buffer, i8* %a_str)
  call i8* @strcat(i8* %buffer, i8* %b_str)
  
  %result = call %Generic* @create(i32 2, i8* %buffer)
  call void @free(i8* %buffer)
  ret %Generic* %result

error:
  call void @panic(i8* getelementptr inbounds ([23 x i8], [23 x i8]* @.error.string_expected, i32 0, i32 0))
  ret %Generic* null
}