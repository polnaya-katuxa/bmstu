#include <syslog.h>
#include <stdlib.h>
#include <fcntl.h>
#include <string.h>
#include <signal.h>
#include <sys/resource.h>
#include <unistd.h>
#include <stdio.h>
#include <stdarg.h>
#include <sys/stat.h>
#include <pthread.h>
#include <errno.h>

// log stream --info --predicate "process == 'a.out'"

sigset_t mask;

void err_quit(const char *fmt, ...)
{
    va_list args;
    
    va_start(args, fmt);
    vfprintf(stderr, fmt, args);
    va_end(args);

    exit(1);
}

int lockfile(int fd)
{
    struct flock fl;

    fl.l_type = F_WRLCK;
    fl.l_start = 0;
    fl.l_whence = SEEK_SET;
    fl.l_len = 0;

    return fcntl(fd, F_SETLK, &fl);
}

void daemonize(const char* cmd)
{
    int i, fd0, fd1, fd2;
    pid_t pid;
    struct rlimit rl;
    struct sigaction sa;

    umask(0);

    if ((pid = fork()) == -1)
        err_quit("%s: fork error", cmd);
    else if (pid > 0)
        exit(0);

    sa.sa_handler = SIG_IGN;
    sigemptyset(&sa.sa_mask); 
    sa.sa_flags = 0; 
    if (sigaction(SIGHUP, &sa, NULL) == -1)
        err_quit("%s: ignore SIGHUP error", cmd);

    if (setsid() == -1)
    	err_quit("%s: setsid error", cmd);

    if (chdir("/") == -1)
        err_quit("%s: change directory error", cmd);

    if (getrlimit(RLIMIT_NOFILE, &rl) == -1)
        err_quit("%s: file limit error", cmd);

    if (rl.rlim_max == RLIM_INFINITY)
        rl.rlim_max = 1024;
    for (int i = 0; i < rl.rlim_max; i++)
        close(i);

    fd0 = open("/dev/null", O_RDWR);
    fd1 = dup(0);
    fd2 = dup(0);

    openlog(cmd, LOG_CONS, LOG_DAEMON); 

    if (fd0 != 0 || fd1 != 1 || fd2 != 2)
    {
        syslog(LOG_ERR, "file descriptors error %d %d %d", fd0, fd1, fd2);
        exit(1);
    }
}

int already_running(void)
{
    int fd;
    char buf[16];
    int perms = S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH;

    fd = open("/var/run/my_daemon.pid", O_RDWR | O_CREAT, perms);
    if (fd == -1)
    {
        syslog(LOG_ERR, "open %s error: %s", "/var/run/my_daemon.pid", strerror(errno));
        exit(1);
    }

    if (lockfile(fd) == -1)
    {
        if (errno == EACCES || errno == EAGAIN)
        {
            close(fd);
            syslog(LOG_ERR, "lock %s error: %s (already locked)", "/var/run/my_daemon.pid", strerror(errno));
            return 1;
        }

        syslog(LOG_ERR, "lock %s error: %s", "/var/run/my_daemon.pid", strerror(errno));
        exit(1);
    }

    if ((ftruncate(fd, 0)) == -1)
    {
        syslog(LOG_ERR, "truncate %s error: %s", "/var/run/my_daemon.pid", strerror(errno));
        exit(1);
    }
    sprintf(buf, "%ld", (long)getpid());
    if ((write(fd, buf, strlen(buf) + 1)) == -1)
    {
        syslog(LOG_ERR, "write to %s error: %s", "/var/run/my_daemon.pid", strerror(errno));
        exit(1);
    }

    return 0;
}

void *thr_fn(void *arg)
{
    int err, signo;

    for (;;)
    {
        err = sigwait(&mask, &signo);
        if (err != 0) {
            syslog(LOG_ERR, "sigwait error");
            exit(1);
        }

        switch (signo)
        {
            case SIGHUP:
                syslog(LOG_INFO, "re-reading configuration file");
                syslog(LOG_INFO, "getlogin: %s", getlogin());
                break;

            case SIGTERM:
                syslog(LOG_INFO, "caught SIGTERM => exit");
                exit(0);	

            default:
                syslog(LOG_WARNING, "unexpected signal %d\n", signo);
        }
    }

    return (void*)0;
}

void logtime()
{
    time_t s;
    struct tm* current_time;
  
    s = time(NULL);
 
    current_time = localtime(&s);
 
    syslog(LOG_INFO, "%02d:%02d:%02d",
           current_time->tm_hour,
           current_time->tm_min,
           current_time->tm_sec);
}

int main(int argc, char* argv[])
{
    char *cmd;
    if ((cmd = strrchr(argv[0], '/')) == NULL)
        cmd = argv[0];
    else
        cmd++;
        
    daemonize(cmd);
    
    if (already_running())
    {
        syslog(LOG_ERR, "daemon already running");
        exit(1);
    }

    struct sigaction sa;
    sa.sa_handler = SIG_DFL;
    sigemptyset(&sa.sa_mask);
    sa.sa_flags = 0;
    if (sigaction(SIGHUP, &sa, NULL) == -1)
    {
        syslog(LOG_ERR, "restore SIGHUP default error");
        exit(1);
    }
        
    sigfillset(&mask);
    int err;
    if ((err = pthread_sigmask(SIG_BLOCK, &mask, NULL)) != 0)
    {
        syslog(LOG_ERR, "SIG_BLOCK error");
        exit(1);
    }

    pthread_t tid;
    if ((err = pthread_create(&tid, NULL, thr_fn, NULL)) != 0)
    {
        syslog(LOG_ERR, "create thread for sig_handler error");
        exit(1);
    }

    while (1)
    {
        sleep(3);
        logtime();
    }
}



