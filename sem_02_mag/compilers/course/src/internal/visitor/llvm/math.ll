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