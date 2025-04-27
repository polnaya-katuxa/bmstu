@.str.int = private unnamed_addr constant [6 x i8] c"%lld\0A\00", align 1
@.str.float = private unnamed_addr constant [4 x i8] c"%f\0A\00", align 1
@.str.string = private unnamed_addr constant [4 x i8] c"%s\0A\00", align 1
@.str.true = private unnamed_addr constant [6 x i8] c"true\0A\00", align 1
@.str.false = private unnamed_addr constant [7 x i8] c"false\0A\00", align 1
@.str.nil = private unnamed_addr constant [5 x i8] c"nil\0A\00", align 1
@.str.brace.left = private unnamed_addr constant [3 x i8] c"{\0A\00", align 1
@.str.brace.right = private unnamed_addr constant [3 x i8] c"}\0A\00", align 1
@.str.tab = private unnamed_addr constant [5 x i8] c"    \00", align 1
@.str.ln = private unnamed_addr constant [2 x i8] c"\0A\00", align 1

define void @print(%Generic* %obj) {
entry:
  %type_ptr = getelementptr inbounds %Generic, %Generic* %obj, i32 0, i32 0
  %type = load i32, i32* %type_ptr
  
  %data_ptr_ptr = getelementptr inbounds %Generic, %Generic* %obj, i32 0, i32 1
  %data_ptr = load i8*, i8** %data_ptr_ptr
  
  switch i32 %type, label %unknown [
    i32 0, label %print_int
    i32 1, label %print_float
    i32 2, label %print_string
    i32 3, label %print_bool
    i32 4, label %print_table
    i32 5, label %print_nil
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

print_nil:
  %nil_fmt = getelementptr inbounds [5 x i8], [5 x i8]* @.str.nil, i32 0, i32 0
  call i32 (i8*, ...) @printf(i8* %nil_fmt)
  ret void

print_table:
  %brace_left_fmt = getelementptr inbounds [3 x i8], [3 x i8]* @.str.brace.left, i32 0, i32 0
  call i32 (i8*, ...) @printf(i8* %brace_left_fmt)
  %tbl = call %LuaTable* @extract_table(%Generic* %obj)
  %valid = icmp ne %LuaTable* %tbl, null
  br i1 %valid, label %proceed, label %unknown

proceed:
  %cur = call %Generic* @create(i32 0, i8* inttoptr (i64 0 to i8*))
	%inc = call %Generic* @create(i32 0, i8* inttoptr (i64 1 to i8*))
	%nil1 = call %Generic* @create_nil()
	%nil2 = call %Generic* @create_nil()
	br label %check_tbl_len

quit:
  %brace_right_fmt = getelementptr inbounds [3 x i8], [3 x i8]* @.str.brace.right, i32 0, i32 0
  call i32 (i8*, ...) @printf(i8* %brace_right_fmt)
	ret void

get_values: ;34
	%cur_key = call %Generic* @lua_table_get_key_at(%Generic* %obj, %Generic* %cur)
	%cur_val = call %Generic* @lua_table_get_value_at(%Generic* %obj, %Generic* %cur)
	call void @copy(%Generic* %cur_key, %Generic* %nil1)
	call void @copy(%Generic* %cur_val, %Generic* %nil2)
	br label %body

check_tbl_len:
	%tbllen = call %Generic* @lua_table_len(%Generic* %obj)
	%lencheck = call %Generic* @ge(%Generic* %cur, %Generic* %tbllen)
	%lennotok = call i1 @check(%Generic* %lencheck)
	br i1 %lennotok, label %quit, label %get_values

inc_iter:
	%iter = call %Generic* @add(%Generic* %cur, %Generic* %inc)
	call void @copy(%Generic* %iter, %Generic* %cur)
	br label %check_tbl_len

body:
  %tab_fmt = getelementptr inbounds [5 x i8], [5 x i8]* @.str.tab, i32 0, i32 0
  call i32 (i8*, ...) @printf(i8* %tab_fmt)
	call void @print(%Generic* %nil1)
  call i32 (i8*, ...) @printf(i8* %tab_fmt)
	call void @print(%Generic* %nil2)
  %ln_fmt = getelementptr inbounds [2 x i8], [2 x i8]* @.str.ln, i32 0, i32 0
  call i32 (i8*, ...) @printf(i8* %ln_fmt)
	br label %inc_iter

unknown:
  call void @panic(i8* getelementptr inbounds ([13 x i8], [13 x i8]* @.error.unknown.type, i32 0, i32 0))
  ret void
}