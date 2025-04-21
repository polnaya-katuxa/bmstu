declare i8* @malloc(i64)
declare void @free(i8*)
declare i32 @strcmp(i8*, i8*)
declare i64 @strlen(i8*)
declare i8* @strcpy(i8*, i8*)
declare i8* @strcat(i8*, i8*)
declare void @llvm.memcpy.p0i8.p0i8.i64(i8*, i8*, i64, i1 immarg)
declare i32 @printf(i8*, ...)
declare double @pow(double, double)
declare void @exit(i64)

@.str.error = private unnamed_addr constant [46 x i8] c"\0A\0A################\20PANIC\20################\0A%s\0A\00", align 1

@.error.error = constant [6 x i8] c"Error\00"

define void @panic(i8* %msg) {
entry:
  %fmt = getelementptr inbounds [46 x i8], [46 x i8]* @.str.error, i32 0, i32 0
  call i32 (i8*, ...) @printf(i8* %fmt, i8* %msg)
  call void @exit(i64 1)
  ret void
}