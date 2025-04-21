@.str.int = private unnamed_addr constant [6 x i8] c"%lld\0A\00", align 1
@.str.float = private unnamed_addr constant [4 x i8] c"%f\0A\00", align 1
@.str.string = private unnamed_addr constant [4 x i8] c"%s\0A\00", align 1
@.str.true = private unnamed_addr constant [6 x i8] c"true\0A\00", align 1
@.str.false = private unnamed_addr constant [7 x i8] c"false\0A\00", align 1
@.str.nil = private unnamed_addr constant [5 x i8] c"nil\0A\00", align 1

; Функция печати
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

unknown:
  ret void
}