@.error.bool_expected = private constant [21 x i8] c"Value should be bool\00"

define %Generic* @not(%Generic* %v) {
entry:
  %v_type_ptr = getelementptr inbounds %Generic, %Generic* %v, i32 0, i32 0
  %v_type = load i32, i32* %v_type_ptr
  
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
  call void @panic(i8* getelementptr inbounds ([21 x i8], [21 x i8]* @.error.bool_expected, i32 0, i32 0))
  ret %Generic* null
}

define %Generic* @and(%Generic* %a, %Generic* %b) {
entry:
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %same_type, label %error_diff

same_type:
  %type_bool = icmp eq i32 %a_type, 3
  br i1 %type_bool, label %and_bool, label %error_bool

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
  
error_bool:
  call void @panic(i8* getelementptr inbounds ([21 x i8], [21 x i8]* @.error.bool_expected, i32 0, i32 0))
  ret %Generic* null

error_diff:
  call void @panic(i8* getelementptr inbounds ([24 x i8], [24 x i8]* @.error.diff.type, i32 0, i32 0))
  ret %Generic* null
}

define %Generic* @or(%Generic* %a, %Generic* %b) {
entry:
  %a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
  %a_type = load i32, i32* %a_type_ptr
  
  %b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
  %b_type = load i32, i32* %b_type_ptr
  
  %type_eq = icmp eq i32 %a_type, %b_type
  br i1 %type_eq, label %same_type, label %error_diff

same_type:
  %type_bool = icmp eq i32 %a_type, 3
  br i1 %type_bool, label %or_bool, label %error_bool

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
  
error_bool:
  call void @panic(i8* getelementptr inbounds ([21 x i8], [21 x i8]* @.error.bool_expected, i32 0, i32 0))
  ret %Generic* null

error_diff:
  call void @panic(i8* getelementptr inbounds ([24 x i8], [24 x i8]* @.error.diff.type, i32 0, i32 0))
  ret %Generic* null
}

define i1 @check(%Generic* %obj) {
entry:
  %is_null = icmp eq %Generic* %obj, null
  br i1 %is_null, label %error_inv, label %check_value

check_value:
  %type_ptr = getelementptr inbounds %Generic, %Generic* %obj, i32 0, i32 0
  %type = load i32, i32* %type_ptr
  
  %data_ptr_ptr = getelementptr inbounds %Generic, %Generic* %obj, i32 0, i32 1
  %data_ptr = load i8*, i8** %data_ptr_ptr
  
  switch i32 %type, label %error_bool [
    i32 3, label %check_bool
  ]

check_bool:
  %bool_ptr = bitcast i8* %data_ptr to i8*
  %bool_val = load i8, i8* %bool_ptr

  %is_true = icmp eq i32 %bool_val, 1
  br i1 %is_true, label %ret_true, label %ret_false

ret_true:
  ret i1 true

ret_false:
  ret i1 false

error_bool:
  call void @panic(i8* getelementptr inbounds ([21 x i8], [21 x i8]* @.error.bool_expected, i32 0, i32 0))
  ret i1 false

error_inv:
  call void @panic(i8* getelementptr inbounds ([17 x i8], [17 x i8]* @.error.invalid.type, i32 0, i32 0))
  ret i1 false
}