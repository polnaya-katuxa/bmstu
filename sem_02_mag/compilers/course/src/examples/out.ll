%Generic = type { i32, i8* }
%LuaTable = type { i32, i32, %TableEntry* }
%TableEntry = type { %Generic*, %Generic*, i1 }

@.str.error = private unnamed_addr constant [46 x i8] c"\0A\0A################ PANIC ################\0A%s\0A\00", align 1
@.error.error = constant [6 x i8] c"Error\00"
@.error.unknown.type = constant [13 x i8] c"Unknown type\00"
@.error.diff.type = constant [24 x i8] c"Different operand types\00"
@.error.invalid.type = constant [17 x i8] c"Invalid operands\00"
@NIL_TYPE = constant i32 5
@.error.null_value = private constant [11 x i8] c"Null value\00"
@.error.bool_expected = private constant [21 x i8] c"Value should be bool\00"
@.error.string_expected = private constant [23 x i8] c"Value should be string\00"
@TABLE_TYPE = constant i32 4
@INITIAL_CAPACITY = constant i32 16
@.error.table_expected = private constant [22 x i8] c"Value should be table\00"
@.error.cannot_extract_table = private constant [21 x i8] c"Cannot extract table\00"
@.error.null_table = private constant [11 x i8] c"Null table\00"
@.error.out_of_table = private constant [13 x i8] c"Out of table\00"
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

declare i8* @malloc(i64 %0)

declare void @free(i8* %0)

declare i32 @strcmp(i8* %0, i8* %1)

declare i64 @strlen(i8* %0)

declare i8* @strcpy(i8* %0, i8* %1)

declare i8* @strcat(i8* %0, i8* %1)

declare void @llvm.memcpy.p0i8.p0i8.i64(i8* %0, i8* %1, i64 %2, i1 immarg %3)

declare i32 @printf(i8* %0, ...)

declare double @pow(double %0, double %1)

declare void @exit(i64 %0)

define void @panic(i8* %msg) {
entry:
	%fmt = getelementptr inbounds [46 x i8], [46 x i8]* @.str.error, i32 0, i32 0
	%0 = call i32 (i8*, ...) @printf(i8* %fmt, i8* %msg)
	call void @exit(i64 1)
	ret void
}

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
	%0 = call i8* @strcpy(i8* %str_space, i8* %value)
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

define %Generic* @neg(%Generic* %v) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %v, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%type_eq_int = icmp eq i32 %a_type, 0
	br i1 %type_eq_int, label %neg_int, label %check_float

check_float:
	%type_eq_float = icmp eq i32 %a_type, 1
	br i1 %type_eq_float, label %neg_float, label %error

neg_float:
	%v_fdata_ptr = getelementptr inbounds %Generic, %Generic* %v, i32 0, i32 1
	%v_fdata = load i8*, i8** %v_fdata_ptr
	%v_fptr = bitcast i8* %v_fdata to double*
	%v_fval = load double, double* %v_fptr
	%v_neg_fval = fneg double %v_fval
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
	%v_val.i8.i_neg = inttoptr i64 %v_val to i8*
	%fresult.i_neg = call %Generic* @create(i32 0, i8* %v_val.i8.i_neg)
	ret %Generic* %fresult.i_neg

error:
	call void @panic(i8* getelementptr inbounds ([17 x i8], [17 x i8]* @.error.invalid.type, i32 0, i32 0))
	ret %Generic* null
}

define %Generic* @add(%Generic* %a, %Generic* %b) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
	%b_type = load i32, i32* %b_type_ptr
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
	%temp.storage.mixed_float_second = alloca double
	store double %fmixedsum.mixed_float_second, double* %temp.storage.mixed_float_second
	%as.i8.mixed_float_second = bitcast double* %temp.storage.mixed_float_second to i8*
	%fresult.mixed_float_second = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_second)
	ret %Generic* %fresult.mixed_float_second

error:
	call void @panic(i8* getelementptr inbounds ([17 x i8], [17 x i8]* @.error.invalid.type, i32 0, i32 0))
	ret %Generic* null
}

define %Generic* @sub(%Generic* %a, %Generic* %b) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
	%b_type = load i32, i32* %b_type_ptr
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
	%temp.storage.mixed_float_second = alloca double
	store double %fmixedsum.mixed_float_second, double* %temp.storage.mixed_float_second
	%as.i8.mixed_float_second = bitcast double* %temp.storage.mixed_float_second to i8*
	%fresult.mixed_float_second = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_second)
	ret %Generic* %fresult.mixed_float_second

error:
	call void @panic(i8* getelementptr inbounds ([17 x i8], [17 x i8]* @.error.invalid.type, i32 0, i32 0))
	ret %Generic* null
}

define %Generic* @mul(%Generic* %a, %Generic* %b) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
	%b_type = load i32, i32* %b_type_ptr
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
	%temp.storage.mixed_float_second = alloca double
	store double %fmixedsum.mixed_float_second, double* %temp.storage.mixed_float_second
	%as.i8.mixed_float_second = bitcast double* %temp.storage.mixed_float_second to i8*
	%fresult.mixed_float_second = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_second)
	ret %Generic* %fresult.mixed_float_second

error:
	call void @panic(i8* getelementptr inbounds ([17 x i8], [17 x i8]* @.error.invalid.type, i32 0, i32 0))
	ret %Generic* null
}

define %Generic* @div(%Generic* %a, %Generic* %b) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
	%b_type = load i32, i32* %b_type_ptr
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
	%temp.storage.mixed_float_second = alloca double
	store double %fmixedsum.mixed_float_second, double* %temp.storage.mixed_float_second
	%as.i8.mixed_float_second = bitcast double* %temp.storage.mixed_float_second to i8*
	%fresult.mixed_float_second = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_second)
	ret %Generic* %fresult.mixed_float_second

error:
	call void @panic(i8* getelementptr inbounds ([17 x i8], [17 x i8]* @.error.invalid.type, i32 0, i32 0))
	ret %Generic* null
}

define %Generic* @mod(%Generic* %a, %Generic* %b) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
	%b_type = load i32, i32* %b_type_ptr
	%type_eq = icmp eq i32 %a_type, %b_type
	br i1 %type_eq, label %same_type, label %error

same_type:
	switch i32 %a_type, label %error [
		i32 0, label %mod_int
		i32 1, label %error
	]

mod_int:
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
	call void @panic(i8* getelementptr inbounds ([17 x i8], [17 x i8]* @.error.invalid.type, i32 0, i32 0))
	ret %Generic* null
}

define %Generic* @rem(%Generic* %a, %Generic* %b) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
	%b_type = load i32, i32* %b_type_ptr
	%type_eq = icmp eq i32 %a_type, %b_type
	br i1 %type_eq, label %same_type, label %error

same_type:
	switch i32 %a_type, label %error [
		i32 0, label %rem_int
		i32 1, label %error
	]

rem_int:
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
	call void @panic(i8* getelementptr inbounds ([17 x i8], [17 x i8]* @.error.invalid.type, i32 0, i32 0))
	ret %Generic* null
}

define %Generic* @power(%Generic* %a, %Generic* %b) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
	%b_type = load i32, i32* %b_type_ptr
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
	%temp.storage.mixed_float_second = alloca double
	store double %fmixedsum.mixed_float_second, double* %temp.storage.mixed_float_second
	%as.i8.mixed_float_second = bitcast double* %temp.storage.mixed_float_second to i8*
	%fresult.mixed_float_second = call %Generic* @create(i32 1, i8* %as.i8.mixed_float_second)
	ret %Generic* %fresult.mixed_float_second

error:
	call void @panic(i8* getelementptr inbounds ([17 x i8], [17 x i8]* @.error.invalid.type, i32 0, i32 0))
	ret %Generic* null
}

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
	%is_true = icmp eq i8 %bool_val, 1
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

declare void @llvm.memset.p0i8.i64(i8* nocapture writeonly %0, i8 %1, i64 %2, i1 immarg %3)

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
	%0 = call i8* @strcpy(i8* %buffer, i8* %a_str)
	%1 = call i8* @strcat(i8* %buffer, i8* %b_str)
	%result = call %Generic* @create(i32 2, i8* %buffer)
	call void @free(i8* %buffer)
	ret %Generic* %result

error:
	call void @panic(i8* getelementptr inbounds ([23 x i8], [23 x i8]* @.error.string_expected, i32 0, i32 0))
	ret %Generic* null
}

define %Generic* @equal(%Generic* %a, %Generic* %b) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
	%b_type = load i32, i32* %b_type_ptr
	%type_eq = icmp eq i32 %a_type, %b_type
	br i1 %type_eq, label %check_value, label %eq_mixed

eq_mixed:
	switch i32 %a_type, label %error [
		i32 0, label %eq_mixed_float_second
		i32 1, label %eq_mixed_float_first
	]

check_value:
	switch i32 %a_type, label %error_unknown [
		i32 0, label %cmp_int
		i32 1, label %cmp_float
		i32 2, label %cmp_str
		i32 3, label %cmp_bool
	]

eq_mixed_float_second:
	%type_eq_mixed_float_second = icmp eq i32 %b_type, 1
	br i1 %type_eq_mixed_float_second, label %proceed_eq_mixed_float_second, label %error

proceed_eq_mixed_float_second:
	%a_fdata_ptr.eq_mixed_float_second = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
	%a_fdata.eq_mixed_float_second = load i8*, i8** %a_fdata_ptr.eq_mixed_float_second
	%a_fptr.eq_mixed_float_second = bitcast i8* %a_fdata.eq_mixed_float_second to i64*
	%a_val.eq_mixed_float_second = load i64, i64* %a_fptr.eq_mixed_float_second
	%a_fval.eq_mixed_float_second = sitofp i64 %a_val.eq_mixed_float_second to double
	%b_fdata_ptr.eq_mixed_float_second = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
	%b_fdata.eq_mixed_float_second = load i8*, i8** %b_fdata_ptr.eq_mixed_float_second
	%b_fptr.eq_mixed_float_second = bitcast i8* %b_fdata.eq_mixed_float_second to double*
	%b_fval.eq_mixed_float_second = load double, double* %b_fptr.eq_mixed_float_second
	%fcmp.eq_mixed_float_second = fcmp oeq double %a_fval.eq_mixed_float_second, %b_fval.eq_mixed_float_second
	%cmp.float.eq_mixed_float_second = inttoptr i1 %fcmp.eq_mixed_float_second to i8*
	%result.float.eq_mixed_float_second = call %Generic* @create(i32 3, i8* %cmp.float.eq_mixed_float_second)
	ret %Generic* %result.float.eq_mixed_float_second

eq_mixed_float_first:
	%type_eq_mixed_float_first = icmp eq i32 %b_type, 0
	br i1 %type_eq_mixed_float_first, label %proceed_eq_mixed_float_first, label %error

proceed_eq_mixed_float_first:
	%a_fdata_ptr.eq_mixed_float_first = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
	%a_fdata.eq_mixed_float_first = load i8*, i8** %a_fdata_ptr.eq_mixed_float_first
	%a_fptr.eq_mixed_float_first = bitcast i8* %a_fdata.eq_mixed_float_first to double*
	%a_fval.eq_mixed_float_first = load double, double* %a_fptr.eq_mixed_float_first
	%b_fdata_ptr.eq_mixed_float_first = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
	%b_fdata.eq_mixed_float_first = load i8*, i8** %b_fdata_ptr.eq_mixed_float_first
	%b_fptr.eq_mixed_float_first = bitcast i8* %b_fdata.eq_mixed_float_first to i64*
	%b_val.eq_mixed_float_first = load i64, i64* %b_fptr.eq_mixed_float_first
	%b_fval.eq_mixed_float_first = sitofp i64 %b_val.eq_mixed_float_first to double
	%fcmp.eq_mixed_float_first = fcmp oeq double %a_fval.eq_mixed_float_first, %b_fval.eq_mixed_float_first
	%cmp.float.eq_mixed_float_first = inttoptr i1 %fcmp.eq_mixed_float_first to i8*
	%result.float.eq_mixed_float_first = call %Generic* @create(i32 3, i8* %cmp.float.eq_mixed_float_first)
	ret %Generic* %result.float.eq_mixed_float_first

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

error_unknown:
	call void @panic(i8* getelementptr inbounds ([13 x i8], [13 x i8]* @.error.unknown.type, i32 0, i32 0))
	ret %Generic* null

error:
	%result.false = call %Generic* @create(i32 3, i8* inttoptr (i8 0 to i8*))
	ret %Generic* %result.false
}

define %Generic* @nequal(%Generic* %a, %Generic* %b) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
	%b_type = load i32, i32* %b_type_ptr
	%type_eq = icmp eq i32 %a_type, %b_type
	br i1 %type_eq, label %check_value, label %neq_mixed

neq_mixed:
	switch i32 %a_type, label %error [
		i32 0, label %neq_mixed_float_second
		i32 1, label %neq_mixed_float_first
	]

check_value:
	switch i32 %a_type, label %error_unknown [
		i32 0, label %cmp_int
		i32 1, label %cmp_float
		i32 2, label %cmp_str
		i32 3, label %cmp_bool
	]

neq_mixed_float_second:
	%type_neq_mixed_float_second = icmp eq i32 %b_type, 1
	br i1 %type_neq_mixed_float_second, label %proceed_neq_mixed_float_second, label %error

proceed_neq_mixed_float_second:
	%a_fdata_ptr.neq_mixed_float_second = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
	%a_fdata.neq_mixed_float_second = load i8*, i8** %a_fdata_ptr.neq_mixed_float_second
	%a_fptr.neq_mixed_float_second = bitcast i8* %a_fdata.neq_mixed_float_second to i64*
	%a_val.neq_mixed_float_second = load i64, i64* %a_fptr.neq_mixed_float_second
	%a_fval.neq_mixed_float_second = sitofp i64 %a_val.neq_mixed_float_second to double
	%b_fdata_ptr.neq_mixed_float_second = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
	%b_fdata.neq_mixed_float_second = load i8*, i8** %b_fdata_ptr.neq_mixed_float_second
	%b_fptr.neq_mixed_float_second = bitcast i8* %b_fdata.neq_mixed_float_second to double*
	%b_fval.neq_mixed_float_second = load double, double* %b_fptr.neq_mixed_float_second
	%fcmp.neq_mixed_float_second = fcmp one double %a_fval.neq_mixed_float_second, %b_fval.neq_mixed_float_second
	%cmp.float.neq_mixed_float_second = inttoptr i1 %fcmp.neq_mixed_float_second to i8*
	%result.float.neq_mixed_float_second = call %Generic* @create(i32 3, i8* %cmp.float.neq_mixed_float_second)
	ret %Generic* %result.float.neq_mixed_float_second

neq_mixed_float_first:
	%type_neq_mixed_float_first = icmp eq i32 %b_type, 0
	br i1 %type_neq_mixed_float_first, label %proceed_neq_mixed_float_first, label %error

proceed_neq_mixed_float_first:
	%a_fdata_ptr.neq_mixed_float_first = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
	%a_fdata.neq_mixed_float_first = load i8*, i8** %a_fdata_ptr.neq_mixed_float_first
	%a_fptr.neq_mixed_float_first = bitcast i8* %a_fdata.neq_mixed_float_first to double*
	%a_fval.neq_mixed_float_first = load double, double* %a_fptr.neq_mixed_float_first
	%b_fdata_ptr.neq_mixed_float_first = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
	%b_fdata.neq_mixed_float_first = load i8*, i8** %b_fdata_ptr.neq_mixed_float_first
	%b_fptr.neq_mixed_float_first = bitcast i8* %b_fdata.neq_mixed_float_first to i64*
	%b_val.neq_mixed_float_first = load i64, i64* %b_fptr.neq_mixed_float_first
	%b_fval.neq_mixed_float_first = sitofp i64 %b_val.neq_mixed_float_first to double
	%fcmp.neq_mixed_float_first = fcmp one double %a_fval.neq_mixed_float_first, %b_fval.neq_mixed_float_first
	%cmp.float.neq_mixed_float_first = inttoptr i1 %fcmp.neq_mixed_float_first to i8*
	%result.float.neq_mixed_float_first = call %Generic* @create(i32 3, i8* %cmp.float.neq_mixed_float_first)
	ret %Generic* %result.float.neq_mixed_float_first

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

error_unknown:
	call void @panic(i8* getelementptr inbounds ([13 x i8], [13 x i8]* @.error.unknown.type, i32 0, i32 0))
	ret %Generic* null

error:
	%result.false = call %Generic* @create(i32 3, i8* inttoptr (i8 0 to i8*))
	ret %Generic* %result.false
}

define %Generic* @gt(%Generic* %a, %Generic* %b) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
	%b_type = load i32, i32* %b_type_ptr
	%type_eq = icmp eq i32 %a_type, %b_type
	br i1 %type_eq, label %check_value, label %gt_mixed

gt_mixed:
	switch i32 %a_type, label %error [
		i32 0, label %gt_mixed_float_second
		i32 1, label %gt_mixed_float_first
	]

check_value:
	switch i32 %a_type, label %error_unknown [
		i32 0, label %cmp_int
		i32 1, label %cmp_float
		i32 2, label %cmp_str
		i32 3, label %cmp_bool
	]

gt_mixed_float_second:
	%type_gt_mixed_float_second = icmp eq i32 %b_type, 1
	br i1 %type_gt_mixed_float_second, label %proceed_gt_mixed_float_second, label %error

proceed_gt_mixed_float_second:
	%a_fdata_ptr.gt_mixed_float_second = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
	%a_fdata.gt_mixed_float_second = load i8*, i8** %a_fdata_ptr.gt_mixed_float_second
	%a_fptr.gt_mixed_float_second = bitcast i8* %a_fdata.gt_mixed_float_second to i64*
	%a_val.gt_mixed_float_second = load i64, i64* %a_fptr.gt_mixed_float_second
	%a_fval.gt_mixed_float_second = sitofp i64 %a_val.gt_mixed_float_second to double
	%b_fdata_ptr.gt_mixed_float_second = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
	%b_fdata.gt_mixed_float_second = load i8*, i8** %b_fdata_ptr.gt_mixed_float_second
	%b_fptr.gt_mixed_float_second = bitcast i8* %b_fdata.gt_mixed_float_second to double*
	%b_fval.gt_mixed_float_second = load double, double* %b_fptr.gt_mixed_float_second
	%fcmp.gt_mixed_float_second = fcmp ogt double %a_fval.gt_mixed_float_second, %b_fval.gt_mixed_float_second
	%cmp.float.gt_mixed_float_second = inttoptr i1 %fcmp.gt_mixed_float_second to i8*
	%result.float.gt_mixed_float_second = call %Generic* @create(i32 3, i8* %cmp.float.gt_mixed_float_second)
	ret %Generic* %result.float.gt_mixed_float_second

gt_mixed_float_first:
	%type_gt_mixed_float_first = icmp eq i32 %b_type, 0
	br i1 %type_gt_mixed_float_first, label %proceed_gt_mixed_float_first, label %error

proceed_gt_mixed_float_first:
	%a_fdata_ptr.gt_mixed_float_first = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
	%a_fdata.gt_mixed_float_first = load i8*, i8** %a_fdata_ptr.gt_mixed_float_first
	%a_fptr.gt_mixed_float_first = bitcast i8* %a_fdata.gt_mixed_float_first to double*
	%a_fval.gt_mixed_float_first = load double, double* %a_fptr.gt_mixed_float_first
	%b_fdata_ptr.gt_mixed_float_first = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
	%b_fdata.gt_mixed_float_first = load i8*, i8** %b_fdata_ptr.gt_mixed_float_first
	%b_fptr.gt_mixed_float_first = bitcast i8* %b_fdata.gt_mixed_float_first to i64*
	%b_val.gt_mixed_float_first = load i64, i64* %b_fptr.gt_mixed_float_first
	%b_fval.gt_mixed_float_first = sitofp i64 %b_val.gt_mixed_float_first to double
	%fcmp.gt_mixed_float_first = fcmp ogt double %a_fval.gt_mixed_float_first, %b_fval.gt_mixed_float_first
	%cmp.float.gt_mixed_float_first = inttoptr i1 %fcmp.gt_mixed_float_first to i8*
	%result.float.gt_mixed_float_first = call %Generic* @create(i32 3, i8* %cmp.float.gt_mixed_float_first)
	ret %Generic* %result.float.gt_mixed_float_first

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

error_unknown:
	call void @panic(i8* getelementptr inbounds ([13 x i8], [13 x i8]* @.error.unknown.type, i32 0, i32 0))
	ret %Generic* null

error:
	%result.false = call %Generic* @create(i32 3, i8* inttoptr (i8 0 to i8*))
	ret %Generic* %result.false
}

define %Generic* @ge(%Generic* %a, %Generic* %b) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
	%b_type = load i32, i32* %b_type_ptr
	%type_eq = icmp eq i32 %a_type, %b_type
	br i1 %type_eq, label %check_value, label %ge_mixed

ge_mixed:
	switch i32 %a_type, label %error [
		i32 0, label %ge_mixed_float_second
		i32 1, label %ge_mixed_float_first
	]

check_value:
	switch i32 %a_type, label %error_unknown [
		i32 0, label %cmp_int
		i32 1, label %cmp_float
		i32 2, label %cmp_str
		i32 3, label %cmp_bool
	]

ge_mixed_float_second:
	%type_ge_mixed_float_second = icmp eq i32 %b_type, 1
	br i1 %type_ge_mixed_float_second, label %proceed_ge_mixed_float_second, label %error

proceed_ge_mixed_float_second:
	%a_fdata_ptr.ge_mixed_float_second = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
	%a_fdata.ge_mixed_float_second = load i8*, i8** %a_fdata_ptr.ge_mixed_float_second
	%a_fptr.ge_mixed_float_second = bitcast i8* %a_fdata.ge_mixed_float_second to i64*
	%a_val.ge_mixed_float_second = load i64, i64* %a_fptr.ge_mixed_float_second
	%a_fval.ge_mixed_float_second = sitofp i64 %a_val.ge_mixed_float_second to double
	%b_fdata_ptr.ge_mixed_float_second = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
	%b_fdata.ge_mixed_float_second = load i8*, i8** %b_fdata_ptr.ge_mixed_float_second
	%b_fptr.ge_mixed_float_second = bitcast i8* %b_fdata.ge_mixed_float_second to double*
	%b_fval.ge_mixed_float_second = load double, double* %b_fptr.ge_mixed_float_second
	%fcmp.ge_mixed_float_second = fcmp oge double %a_fval.ge_mixed_float_second, %b_fval.ge_mixed_float_second
	%cmp.float.ge_mixed_float_second = inttoptr i1 %fcmp.ge_mixed_float_second to i8*
	%result.float.ge_mixed_float_second = call %Generic* @create(i32 3, i8* %cmp.float.ge_mixed_float_second)
	ret %Generic* %result.float.ge_mixed_float_second

ge_mixed_float_first:
	%type_ge_mixed_float_first = icmp eq i32 %b_type, 0
	br i1 %type_ge_mixed_float_first, label %proceed_ge_mixed_float_first, label %error

proceed_ge_mixed_float_first:
	%a_fdata_ptr.ge_mixed_float_first = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
	%a_fdata.ge_mixed_float_first = load i8*, i8** %a_fdata_ptr.ge_mixed_float_first
	%a_fptr.ge_mixed_float_first = bitcast i8* %a_fdata.ge_mixed_float_first to double*
	%a_fval.ge_mixed_float_first = load double, double* %a_fptr.ge_mixed_float_first
	%b_fdata_ptr.ge_mixed_float_first = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
	%b_fdata.ge_mixed_float_first = load i8*, i8** %b_fdata_ptr.ge_mixed_float_first
	%b_fptr.ge_mixed_float_first = bitcast i8* %b_fdata.ge_mixed_float_first to i64*
	%b_val.ge_mixed_float_first = load i64, i64* %b_fptr.ge_mixed_float_first
	%b_fval.ge_mixed_float_first = sitofp i64 %b_val.ge_mixed_float_first to double
	%fcmp.ge_mixed_float_first = fcmp oge double %a_fval.ge_mixed_float_first, %b_fval.ge_mixed_float_first
	%cmp.float.ge_mixed_float_first = inttoptr i1 %fcmp.ge_mixed_float_first to i8*
	%result.float.ge_mixed_float_first = call %Generic* @create(i32 3, i8* %cmp.float.ge_mixed_float_first)
	ret %Generic* %result.float.ge_mixed_float_first

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
		i1 false, label %equal_string
		i1 true, label %geq_string
	]

equal_string:
	switch i1 %res.string.equal, label %error [
		i1 false, label %ngeq_string
		i1 true, label %geq_string
	]

geq_string:
	%cmp.string.geq = inttoptr i1 true to i8*
	%result.string.geq = call %Generic* @create(i32 3, i8* %cmp.string.geq)
	ret %Generic* %result.string.geq

ngeq_string:
	%cmp.string.ngeq = inttoptr i1 false to i8*
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

error_unknown:
	call void @panic(i8* getelementptr inbounds ([13 x i8], [13 x i8]* @.error.unknown.type, i32 0, i32 0))
	ret %Generic* null

error:
	%result.false = call %Generic* @create(i32 3, i8* inttoptr (i8 0 to i8*))
	ret %Generic* %result.false
}

define %Generic* @lt(%Generic* %a, %Generic* %b) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
	%b_type = load i32, i32* %b_type_ptr
	%type_eq = icmp eq i32 %a_type, %b_type
	br i1 %type_eq, label %check_value, label %lt_mixed

check_value:
	switch i32 %a_type, label %error_unknown [
		i32 0, label %cmp_int
		i32 1, label %cmp_float
		i32 2, label %cmp_str
		i32 3, label %cmp_bool
	]

lt_mixed:
	switch i32 %a_type, label %error [
		i32 0, label %lt_mixed_float_second
		i32 1, label %lt_mixed_float_first
	]

lt_mixed_float_second:
	%type_lt_mixed_float_second = icmp eq i32 %b_type, 1
	br i1 %type_lt_mixed_float_second, label %proceed_lt_mixed_float_second, label %error

proceed_lt_mixed_float_second:
	%a_fdata_ptr.lt_mixed_float_second = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
	%a_fdata.lt_mixed_float_second = load i8*, i8** %a_fdata_ptr.lt_mixed_float_second
	%a_fptr.lt_mixed_float_second = bitcast i8* %a_fdata.lt_mixed_float_second to i64*
	%a_val.lt_mixed_float_second = load i64, i64* %a_fptr.lt_mixed_float_second
	%a_fval.lt_mixed_float_second = sitofp i64 %a_val.lt_mixed_float_second to double
	%b_fdata_ptr.lt_mixed_float_second = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
	%b_fdata.lt_mixed_float_second = load i8*, i8** %b_fdata_ptr.lt_mixed_float_second
	%b_fptr.lt_mixed_float_second = bitcast i8* %b_fdata.lt_mixed_float_second to double*
	%b_fval.lt_mixed_float_second = load double, double* %b_fptr.lt_mixed_float_second
	%fcmp.lt_mixed_float_second = fcmp olt double %a_fval.lt_mixed_float_second, %b_fval.lt_mixed_float_second
	%cmp.float.lt_mixed_float_second = inttoptr i1 %fcmp.lt_mixed_float_second to i8*
	%result.float.lt_mixed_float_second = call %Generic* @create(i32 3, i8* %cmp.float.lt_mixed_float_second)
	ret %Generic* %result.float.lt_mixed_float_second

lt_mixed_float_first:
	%type_lt_mixed_float_first = icmp eq i32 %b_type, 0
	br i1 %type_lt_mixed_float_first, label %proceed_lt_mixed_float_first, label %error

proceed_lt_mixed_float_first:
	%a_fdata_ptr.lt_mixed_float_first = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
	%a_fdata.lt_mixed_float_first = load i8*, i8** %a_fdata_ptr.lt_mixed_float_first
	%a_fptr.lt_mixed_float_first = bitcast i8* %a_fdata.lt_mixed_float_first to double*
	%a_fval.lt_mixed_float_first = load double, double* %a_fptr.lt_mixed_float_first
	%b_fdata_ptr.lt_mixed_float_first = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
	%b_fdata.lt_mixed_float_first = load i8*, i8** %b_fdata_ptr.lt_mixed_float_first
	%b_fptr.lt_mixed_float_first = bitcast i8* %b_fdata.lt_mixed_float_first to i64*
	%b_val.lt_mixed_float_first = load i64, i64* %b_fptr.lt_mixed_float_first
	%b_fval.lt_mixed_float_first = sitofp i64 %b_val.lt_mixed_float_first to double
	%fcmp.lt_mixed_float_first = fcmp olt double %a_fval.lt_mixed_float_first, %b_fval.lt_mixed_float_first
	%cmp.float.lt_mixed_float_first = inttoptr i1 %fcmp.lt_mixed_float_first to i8*
	%result.float.lt_mixed_float_first = call %Generic* @create(i32 3, i8* %cmp.float.lt_mixed_float_first)
	ret %Generic* %result.float.lt_mixed_float_first

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

error_unknown:
	call void @panic(i8* getelementptr inbounds ([13 x i8], [13 x i8]* @.error.unknown.type, i32 0, i32 0))
	ret %Generic* null

error:
	%result.false = call %Generic* @create(i32 3, i8* inttoptr (i8 0 to i8*))
	ret %Generic* %result.false
}

define %Generic* @le(%Generic* %a, %Generic* %b) {
entry:
	%a_type_ptr = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 0
	%a_type = load i32, i32* %a_type_ptr
	%b_type_ptr = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 0
	%b_type = load i32, i32* %b_type_ptr
	%type_eq = icmp eq i32 %a_type, %b_type
	br i1 %type_eq, label %check_value, label %le_mixed

check_value:
	switch i32 %a_type, label %error_unknown [
		i32 0, label %cmp_int
		i32 1, label %cmp_float
		i32 2, label %cmp_str
		i32 3, label %cmp_bool
	]

le_mixed:
	switch i32 %a_type, label %error [
		i32 0, label %le_mixed_float_second
		i32 1, label %le_mixed_float_first
	]

le_mixed_float_second:
	%type_le_mixed_float_second = icmp eq i32 %b_type, 1
	br i1 %type_le_mixed_float_second, label %proceed_le_mixed_float_second, label %error

proceed_le_mixed_float_second:
	%a_fdata_ptr.le_mixed_float_second = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
	%a_fdata.le_mixed_float_second = load i8*, i8** %a_fdata_ptr.le_mixed_float_second
	%a_fptr.le_mixed_float_second = bitcast i8* %a_fdata.le_mixed_float_second to i64*
	%a_val.le_mixed_float_second = load i64, i64* %a_fptr.le_mixed_float_second
	%a_fval.le_mixed_float_second = sitofp i64 %a_val.le_mixed_float_second to double
	%b_fdata_ptr.le_mixed_float_second = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
	%b_fdata.le_mixed_float_second = load i8*, i8** %b_fdata_ptr.le_mixed_float_second
	%b_fptr.le_mixed_float_second = bitcast i8* %b_fdata.le_mixed_float_second to double*
	%b_fval.le_mixed_float_second = load double, double* %b_fptr.le_mixed_float_second
	%fcmp.le_mixed_float_second = fcmp ole double %a_fval.le_mixed_float_second, %b_fval.le_mixed_float_second
	%cmp.float.le_mixed_float_second = inttoptr i1 %fcmp.le_mixed_float_second to i8*
	%result.float.le_mixed_float_second = call %Generic* @create(i32 3, i8* %cmp.float.le_mixed_float_second)
	ret %Generic* %result.float.le_mixed_float_second

le_mixed_float_first:
	%type_le_mixed_float_first = icmp eq i32 %b_type, 0
	br i1 %type_le_mixed_float_first, label %proceed_le_mixed_float_first, label %error

proceed_le_mixed_float_first:
	%a_fdata_ptr.le_mixed_float_first = getelementptr inbounds %Generic, %Generic* %a, i32 0, i32 1
	%a_fdata.le_mixed_float_first = load i8*, i8** %a_fdata_ptr.le_mixed_float_first
	%a_fptr.le_mixed_float_first = bitcast i8* %a_fdata.le_mixed_float_first to double*
	%a_fval.le_mixed_float_first = load double, double* %a_fptr.le_mixed_float_first
	%b_fdata_ptr.le_mixed_float_first = getelementptr inbounds %Generic, %Generic* %b, i32 0, i32 1
	%b_fdata.le_mixed_float_first = load i8*, i8** %b_fdata_ptr.le_mixed_float_first
	%b_fptr.le_mixed_float_first = bitcast i8* %b_fdata.le_mixed_float_first to i64*
	%b_val.le_mixed_float_first = load i64, i64* %b_fptr.le_mixed_float_first
	%b_fval.le_mixed_float_first = sitofp i64 %b_val.le_mixed_float_first to double
	%fcmp.le_mixed_float_first = fcmp ole double %a_fval.le_mixed_float_first, %b_fval.le_mixed_float_first
	%cmp.float.le_mixed_float_first = inttoptr i1 %fcmp.le_mixed_float_first to i8*
	%result.float.le_mixed_float_first = call %Generic* @create(i32 3, i8* %cmp.float.le_mixed_float_first)
	ret %Generic* %result.float.le_mixed_float_first

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
		i1 false, label %equal_string
		i1 true, label %leq_string
	]

equal_string:
	switch i1 %res.string.equal, label %error [
		i1 false, label %nleq_string
		i1 true, label %leq_string
	]

leq_string:
	%cmp.string.leq = inttoptr i1 true to i8*
	%result.string.leq = call %Generic* @create(i32 3, i8* %cmp.string.leq)
	ret %Generic* %result.string.leq

nleq_string:
	%cmp.string.nleq = inttoptr i1 false to i8*
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

error_unknown:
	call void @panic(i8* getelementptr inbounds ([13 x i8], [13 x i8]* @.error.unknown.type, i32 0, i32 0))
	ret %Generic* null

error:
	%result.false = call %Generic* @create(i32 3, i8* inttoptr (i8 0 to i8*))
	ret %Generic* %result.false
}

define %Generic* @lua_table_new() {
entry:
	%capacity = load i32, i32* @INITIAL_CAPACITY
	%entry_size = ptrtoint %TableEntry* getelementptr inbounds (%TableEntry, %TableEntry* null, i32 1) to i64
	%capacity_sext = sext i32 %capacity to i64
	%entries_size = mul i64 %entry_size, %capacity_sext
	%entries = call i8* @malloc(i64 %entries_size)
	%entries_ptr = bitcast i8* %entries to %TableEntry*
	br label %init_loop

init_loop:
	%i = phi i32 [ 0, %entry ], [ %next_i, %init_loop ]
	%entry_ptr = getelementptr inbounds %TableEntry, %TableEntry* %entries_ptr, i32 %i
	%flag_ptr = getelementptr inbounds %TableEntry, %TableEntry* %entry_ptr, i32 0, i32 2
	store i1 false, i1* %flag_ptr, align 1
	%next_i = add i32 %i, 1
	%done = icmp eq i32 %next_i, %capacity
	br i1 %done, label %create_table, label %init_loop

create_table:
	%table_size = ptrtoint %LuaTable* getelementptr inbounds (%LuaTable, %LuaTable* null, i32 1) to i64
	%table = call i8* @malloc(i64 %table_size)
	%null_table = icmp eq i8* %table, null
	br i1 %null_table, label %error, label %continue

continue:
	%table_ptr = bitcast i8* %table to %LuaTable*
	%size_ptr = getelementptr inbounds %LuaTable, %LuaTable* %table_ptr, i32 0, i32 0
	store i32 0, i32* %size_ptr, align 4
	%cap_ptr = getelementptr inbounds %LuaTable, %LuaTable* %table_ptr, i32 0, i32 1
	store i32 %capacity, i32* %cap_ptr, align 4
	%entries_field = getelementptr inbounds %LuaTable, %LuaTable* %table_ptr, i32 0, i32 2
	store %TableEntry* %entries_ptr, %TableEntry** %entries_field, align 8
	%generic = call %Generic* @create(i32 4, i8* %table)
	ret %Generic* %generic

error:
	call void @panic(i8* getelementptr inbounds ([11 x i8], [11 x i8]* @.error.null_table, i32 0, i32 0))
	ret %Generic* null
}

define void @lua_table_set(%Generic* %table, %Generic* %key, %Generic* %value) {
entry:
	%tbl = call %LuaTable* @extract_table(%Generic* %table)
	%is_valid = icmp ne %LuaTable* %tbl, null
	br i1 %is_valid, label %proceed, label %error

proceed:
	%key_copy = call %Generic* @create_nil()
	call void @copy(%Generic* %key, %Generic* %key_copy)
	%value_copy = call %Generic* @create_nil()
	call void @copy(%Generic* %value, %Generic* %value_copy)
	%entries_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 2
	%entries = load %TableEntry*, %TableEntry** %entries_ptr, align 8
	%capacity_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 1
	%capacity = load i32, i32* %capacity_ptr, align 4
	%size_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 0
	%size = load i32, i32* %size_ptr, align 4
	br label %search_loop

search_loop:
	%i = phi i32 [ 0, %proceed ], [ %next_i, %next ]
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
	call void @panic(i8* getelementptr inbounds ([11 x i8], [11 x i8]* @.error.null_table, i32 0, i32 0))
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
	%i = phi i32 [ 0, %proceed ], [ %next_i, %next ]
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
	call void @panic(i8* getelementptr inbounds ([11 x i8], [11 x i8]* @.error.null_table, i32 0, i32 0))
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

define void @resize_table(%LuaTable* %tbl) {
entry:
	%old_capacity_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 1
	%old_capacity = load i32, i32* %old_capacity_ptr, align 4
	%new_capacity = mul i32 %old_capacity, 2
	%entry_size = ptrtoint %TableEntry* getelementptr inbounds (%TableEntry, %TableEntry* null, i32 1) to i64
	%new_capacity_sext = sext i32 %new_capacity to i64
	%new_entries_size = mul i64 %entry_size, %new_capacity_sext
	%new_entries = call i8* @malloc(i64 %new_entries_size)
	%new_entries_ptr = bitcast i8* %new_entries to %TableEntry*
	%old_entries_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 2
	%old_entries = load %TableEntry*, %TableEntry** %old_entries_ptr, align 8
	br label %copy_loop

copy_loop:
	%i = phi i32 [ 0, %entry ], [ %next_i, %next ]
	%current_old = getelementptr inbounds %TableEntry, %TableEntry* %old_entries, i32 %i
	%current_new = getelementptr inbounds %TableEntry, %TableEntry* %new_entries_ptr, i32 %i
	%flag_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current_old, i32 0, i32 2
	%flag = load i1, i1* %flag_ptr, align 1
	%new_flag_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current_new, i32 0, i32 2
	store i1 %flag, i1* %new_flag_ptr, align 1
	br i1 %flag, label %copy_data, label %next

copy_data:
	%old_key_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current_old, i32 0, i32 0
	%old_key = load %Generic*, %Generic** %old_key_ptr, align 8
	%new_key_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current_new, i32 0, i32 0
	%key_copy = call %Generic* @create_nil()
	call void @copy(%Generic* %old_key, %Generic* %key_copy)
	store %Generic* %key_copy, %Generic** %new_key_ptr, align 8
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
	%old_entries_i8 = bitcast %TableEntry* %old_entries to i8*
	call void @free(i8* %old_entries_i8)
	%new_entries_field = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 2
	store %TableEntry* %new_entries_ptr, %TableEntry** %new_entries_field, align 8
	%new_capacity_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 1
	store i32 %new_capacity, i32* %new_capacity_ptr, align 4
	ret void
}

define %Generic* @lua_table_len(%Generic* %table) {
entry:
	%tbl = call %LuaTable* @extract_table(%Generic* %table)
	%is_valid = icmp ne %LuaTable* %tbl, null
	br i1 %is_valid, label %valid, label %error

valid:
	%size_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 0
	%size = load i32, i32* %size_ptr, align 4
	%size_i64 = sext i32 %size to i64
	%size_i8 = inttoptr i64 %size_i64 to i8*
	%size_gen = call %Generic* @create(i32 0, i8* %size_i8)
	ret %Generic* %size_gen

error:
	call void @panic(i8* getelementptr inbounds ([11 x i8], [11 x i8]* @.error.null_table, i32 0, i32 0))
	%nil = call %Generic* @create_nil()
	ret %Generic* %nil
}

define %Generic* @lua_table_get_key_at(%Generic* %table, %Generic* %ind) {
entry:
	%v_data_ptr_int = getelementptr inbounds %Generic, %Generic* %ind, i32 0, i32 1
	%v_data_int = load i8*, i8** %v_data_ptr_int
	%v_ptr_int = bitcast i8* %v_data_int to i64*
	%v_val_int = load i64, i64* %v_ptr_int
	%index = trunc i64 %v_val_int to i32
	%tbl = call %LuaTable* @extract_table(%Generic* %table)
	%is_valid = icmp ne %LuaTable* %tbl, null
	br i1 %is_valid, label %proceed, label %error_null

proceed:
	%entries_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 2
	%entries = load %TableEntry*, %TableEntry** %entries_ptr, align 8
	%capacity_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 1
	%capacity = load i32, i32* %capacity_ptr, align 4
	br label %search_loop

search_loop:
	%i = phi i32 [ 0, %proceed ], [ %next_i, %next ]
	%count = phi i32 [ 0, %proceed ], [ %new_count, %next ]
	%current = getelementptr inbounds %TableEntry, %TableEntry* %entries, i32 %i
	%occupied_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current, i32 0, i32 2
	%occupied = load i1, i1* %occupied_ptr, align 1
	%found = icmp eq i32 %count, %index
	br i1 %occupied, label %check_index, label %next

check_index:
	br i1 %found, label %extract_key, label %increment_count

increment_count:
	%new_count = add i32 %count, 1
	br label %next

extract_key:
	%key_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current, i32 0, i32 0
	%key = load %Generic*, %Generic** %key_ptr, align 8
	%key_copy = call %Generic* @create_nil()
	call void @copy(%Generic* %key, %Generic* %key_copy)
	ret %Generic* %key_copy

next:
	%next_i = add i32 %i, 1
	%in_bounds = icmp ult i32 %next_i, %capacity
	br i1 %in_bounds, label %search_loop, label %error

error_null:
	call void @panic(i8* getelementptr inbounds ([11 x i8], [11 x i8]* @.error.null_table, i32 0, i32 0))
	%nil1 = call %Generic* @create_nil()
	ret %Generic* %nil1

error:
	call void @panic(i8* getelementptr inbounds ([13 x i8], [13 x i8]* @.error.out_of_table, i32 0, i32 0))
	%nil2 = call %Generic* @create_nil()
	ret %Generic* %nil2
}

define %Generic* @lua_table_get_value_at(%Generic* %table, %Generic* %ind) {
entry:
	%v_data_ptr_int = getelementptr inbounds %Generic, %Generic* %ind, i32 0, i32 1
	%v_data_int = load i8*, i8** %v_data_ptr_int
	%v_ptr_int = bitcast i8* %v_data_int to i64*
	%v_val_int = load i64, i64* %v_ptr_int
	%index = trunc i64 %v_val_int to i32
	%tbl = call %LuaTable* @extract_table(%Generic* %table)
	%is_valid = icmp ne %LuaTable* %tbl, null
	br i1 %is_valid, label %proceed, label %error_null

proceed:
	%entries_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 2
	%entries = load %TableEntry*, %TableEntry** %entries_ptr, align 8
	%capacity_ptr = getelementptr inbounds %LuaTable, %LuaTable* %tbl, i32 0, i32 1
	%capacity = load i32, i32* %capacity_ptr, align 4
	br label %search_loop

search_loop:
	%i = phi i32 [ 0, %proceed ], [ %next_i, %next ]
	%count = phi i32 [ 0, %proceed ], [ %new_count, %next ]
	%current = getelementptr inbounds %TableEntry, %TableEntry* %entries, i32 %i
	%occupied_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current, i32 0, i32 2
	%occupied = load i1, i1* %occupied_ptr, align 1
	%found = icmp eq i32 %count, %index
	br i1 %occupied, label %check_index, label %next

check_index:
	br i1 %found, label %extract_value, label %increment_count

increment_count:
	%new_count = add i32 %count, 1
	br label %next

extract_value:
	%value_ptr = getelementptr inbounds %TableEntry, %TableEntry* %current, i32 0, i32 1
	%value = load %Generic*, %Generic** %value_ptr, align 8
	%value_copy = call %Generic* @create_nil()
	call void @copy(%Generic* %value, %Generic* %value_copy)
	ret %Generic* %value_copy

next:
	%next_i = add i32 %i, 1
	%in_bounds = icmp ult i32 %next_i, %capacity
	br i1 %in_bounds, label %search_loop, label %error

error_null:
	call void @panic(i8* getelementptr inbounds ([11 x i8], [11 x i8]* @.error.null_table, i32 0, i32 0))
	%nil1 = call %Generic* @create_nil()
	ret %Generic* %nil1

error:
	call void @panic(i8* getelementptr inbounds ([13 x i8], [13 x i8]* @.error.out_of_table, i32 0, i32 0))
	%nil2 = call %Generic* @create_nil()
	ret %Generic* %nil2
}

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
	%0 = call i32 (i8*, ...) @printf(i8* %int_fmt, i64 %int_val)
	ret void

print_float:
	%float_ptr = bitcast i8* %data_ptr to double*
	%float_val = load double, double* %float_ptr
	%float_fmt = getelementptr inbounds [4 x i8], [4 x i8]* @.str.float, i32 0, i32 0
	%1 = call i32 (i8*, ...) @printf(i8* %float_fmt, double %float_val)
	ret void

print_string:
	%string_val = bitcast i8* %data_ptr to i8*
	%string_fmt = getelementptr inbounds [4 x i8], [4 x i8]* @.str.string, i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %string_fmt, i8* %string_val)
	ret void

print_bool:
	%bool_ptr = bitcast i8* %data_ptr to i8*
	%bool_val = load i8, i8* %bool_ptr
	%is_true = icmp eq i8 %bool_val, 1
	br i1 %is_true, label %print_true, label %print_false

print_true:
	%true_fmt = getelementptr inbounds [6 x i8], [6 x i8]* @.str.true, i32 0, i32 0
	%3 = call i32 (i8*, ...) @printf(i8* %true_fmt)
	ret void

print_false:
	%false_fmt = getelementptr inbounds [7 x i8], [7 x i8]* @.str.false, i32 0, i32 0
	%4 = call i32 (i8*, ...) @printf(i8* %false_fmt)
	ret void

print_nil:
	%nil_fmt = getelementptr inbounds [5 x i8], [5 x i8]* @.str.nil, i32 0, i32 0
	%5 = call i32 (i8*, ...) @printf(i8* %nil_fmt)
	ret void

print_table:
	%brace_left_fmt = getelementptr inbounds [3 x i8], [3 x i8]* @.str.brace.left, i32 0, i32 0
	%6 = call i32 (i8*, ...) @printf(i8* %brace_left_fmt)
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
	%7 = call i32 (i8*, ...) @printf(i8* %brace_right_fmt)
	ret void

get_values:
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
	%8 = call i32 (i8*, ...) @printf(i8* %tab_fmt)
	call void @print(%Generic* %nil1)
	%9 = call i32 (i8*, ...) @printf(i8* %tab_fmt)
	call void @print(%Generic* %nil2)
	%ln_fmt = getelementptr inbounds [2 x i8], [2 x i8]* @.str.ln, i32 0, i32 0
	%10 = call i32 (i8*, ...) @printf(i8* %ln_fmt)
	br label %inc_iter

unknown:
	call void @panic(i8* getelementptr inbounds ([13 x i8], [13 x i8]* @.error.unknown.type, i32 0, i32 0))
	ret void
}

define i64 @main() {
0:
	%1 = call %Generic* @create(i32 0, i8* inttoptr (i64 7 to i8*))
	%2 = call %Generic* @fibonacci(%Generic* %1)
	call void @print(%Generic* %2)
	ret i64 0
}

define %Generic* @fibonacci(%Generic* %n) {
0:
	%1 = call %Generic* @create(i32 0, i8* inttoptr (i64 0 to i8*))
	%2 = call %Generic* @create(i32 0, i8* inttoptr (i64 1 to i8*))
	%3 = call %Generic* @create(i32 0, i8* inttoptr (i64 0 to i8*))
	%4 = call %Generic* @create(i32 0, i8* inttoptr (i64 0 to i8*))
	%5 = call %Generic* @create(i32 0, i8* inttoptr (i64 1 to i8*))
	br label %6

6:
	%7 = call %Generic* @lt(%Generic* %5, %Generic* %3)
	%8 = call i1 @check(%Generic* %7)
	br i1 %8, label %17, label %14

9:
	ret %Generic* %1

10:
	%11 = call %Generic* @add(%Generic* %4, %Generic* %5)
	call void @copy(%Generic* %11, %Generic* %4)
	br label %6

12:
	%13 = call %Generic* @add(%Generic* %1, %Generic* %2)
	call void @copy(%Generic* %2, %Generic* %1)
	call void @copy(%Generic* %13, %Generic* %2)
	br label %10

14:
	%15 = call %Generic* @lt(%Generic* %4, %Generic* %n)
	%16 = call i1 @check(%Generic* %15)
	br i1 %16, label %12, label %9

17:
	%18 = call %Generic* @gt(%Generic* %4, %Generic* %n)
	%19 = call i1 @check(%Generic* %18)
	br i1 %19, label %12, label %9
}
