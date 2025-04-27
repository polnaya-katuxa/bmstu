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
  
error_unknown:
  call void @panic(i8* getelementptr inbounds ([13 x i8], [13 x i8]* @.error.unknown.type, i32 0, i32 0))
  ret %Generic* null

error:
  %result.false = call %Generic* @create(i32 3, i8* inttoptr (i8 0 to i8*))
  ret %Generic* %result.false
}