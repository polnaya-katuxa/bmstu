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