#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/string.h>
#include <linux/vmalloc.h>
#include <asm/uaccess.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Aleksey Knyazhev");

#define DIRNAME "seqfiles"
#define FILENAME "seqfile"
#define SYMLINK "seqfile_link"
#define FILEPATH DIRNAME "/" FILENAME

#define BUF_SIZE PAGE_SIZE

static struct proc_dir_entry *fortune_file = NULL;
static struct proc_dir_entry *fortune_dir = NULL;
static struct proc_dir_entry *fortune_link = NULL;

static char *cookie_buffer = NULL;  
static int write_index = 0;  
static int read_index = 0;  

static char tmp[BUF_SIZE];

ssize_t seq_file_write(struct file *filep, const char __user *buf, size_t len, loff_t *offp) 
{
	printk(KERN_INFO "+ seq_file: write");
	if(len > BUF_SIZE - write_index + 1) 
	{
		printk(KERN_ERR "+ seq_file: cookie_buffer overflow");
		return -ENOSPC;
	}
	if(copy_from_user(&cookie_buffer[write_index], buf, len)) 
	{
		printk(KERN_ERR "+ seq_file: copy from user error");
		return -EFAULT;
	}
	write_index += len;
	cookie_buffer[write_index-1] = '\0';
	return len;
}

static int seq_file_show(struct seq_file *m, void *v) 
{
	printk(KERN_INFO "+ seq_file: show");
	if (!write_index)
		return 0;
	if (read_index >= write_index) 
	{
		read_index = 0;
	}
	int len = snprintf(tmp, BUF_SIZE, "%s\n", &cookie_buffer[read_index]);
	seq_printf(m, "%s", tmp);
	read_index += len;
	return 0;
} 

static void *seq_file_start(struct seq_file *m, loff_t *pos)
{
	printk(KERN_INFO "+ seq_file: start");
	static unsigned long counter = 0;
	if (!*pos) {
		return &counter;
	}
	else {
		*pos = 0;
		return NULL;
	}
}

static void *seq_file_next(struct seq_file *m, void *v, loff_t *pos) 
{
	printk(KERN_INFO "+ seq_file: next");
	unsigned long *tmp = (unsigned long *) v;
	(*tmp)++;
	(*pos)++;
	return NULL;
}

static void seq_file_stop(struct seq_file *m, void *v) 
{
	printk(KERN_INFO "+ seq_file: stop");
}

static struct seq_operations seq_file_ops = {
	.start = seq_file_start,
	.next = seq_file_next,
	.stop = seq_file_stop,
	.show = seq_file_show
};

static int seq_file_open(struct inode *i, struct file * f) 
{
	printk(KERN_DEBUG "+ seq_file: open seq_file");
	return seq_open(f, &seq_file_ops);
}

static struct proc_ops fops = {
	.proc_open = seq_file_open,
	.proc_read = seq_read,
	.proc_write = seq_file_write,
	.proc_lseek = seq_lseek,
	.proc_release = seq_release, 
};

void freemem(void) 
{
	if (cookie_buffer)
		vfree(cookie_buffer);
	if (fortune_link)
		remove_proc_entry(SYMLINK, NULL);
	if (fortune_file)
		remove_proc_entry(FILENAME, fortune_dir);
	if (fortune_dir)
		remove_proc_entry(DIRNAME, NULL);
}

static int __init init_seq_file_module(void) 
{
	if (!(cookie_buffer = vmalloc(BUF_SIZE))) 
	{
		freemem();
		printk(KERN_ERR "+ seq_file: error during vmalloc");
		return -ENOMEM;
	}
	memset(cookie_buffer, 0, BUF_SIZE);
	if (!(fortune_dir = proc_mkdir(DIRNAME, NULL))) 
	{
		freemem();
		printk(KERN_ERR "+ seq_file: error during directory creation");
		return -ENOMEM;
	} 
	else if (!(fortune_file = proc_create(FILENAME, 0666, fortune_dir, &fops))) 
	{
		freemem();
		printk(KERN_ERR "+ seq_file: error during file creation");
		return -ENOMEM;
	} 
	else if (!(fortune_link = proc_symlink(SYMLINK, NULL, FILEPATH))) 
	{
		freemem();
		printk(KERN_ERR "+ seq_file: error during symlink creation");
		return -ENOMEM;
	}
	write_index = 0;
	read_index = 0;
	printk(KERN_INFO "+ module loaded");
	return 0;
}

static void __exit exit_seq_file_module(void) 
{
	freemem();
	printk(KERN_INFO "+ seq_file: unloaded");
}

module_init(init_seq_file_module);
module_exit(exit_seq_file_module);