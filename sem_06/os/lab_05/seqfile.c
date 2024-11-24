#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/string.h>
#include <linux/vmalloc.h>
//#include <asm/uaccess.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Karpova Ekaterina");

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

ssize_t my_write(struct file *file, const char __user *buf, size_t len, loff_t *offp) 
{
	printk(KERN_INFO "+ seqfile: called write");
	if(len > BUF_SIZE - write_index + 1) 
	{
		printk(KERN_ERR "+ seqfile: cookie_buffer overflow");
		return -ENOSPC;
	}
	if(copy_from_user(&cookie_buffer[write_index], buf, len)) 
	{
		printk(KERN_ERR "+ seqfile: copy from user error");
		return -EFAULT;
	}
	write_index += len;
	cookie_buffer[write_index-1] = '\0';
	return len;
}

static int seqfile_show(struct seq_file *m, void *v) 
{
	printk(KERN_INFO "+ seqfile: called show");
	//if (read_index >= write_index)
		//return 0;
	int len = snprintf(tmp, BUF_SIZE, "%s\n", &cookie_buffer[read_index]);
	seq_printf(m, "%s", tmp);
	read_index += len;
	return 0;
} 

ssize_t seqfile_read(struct file *file, char __user *buf, size_t count, loff_t *offp) 
{
	printk(KERN_INFO "+ seqfile: called read");
	return seq_read(file, buf, count, offp);
}

int seqfile_open(struct inode *inode, struct file *file) 
{
	printk(KERN_INFO "+ seqfile: called open");
	return single_open(file, seqfile_show, NULL);
}

int seqfile_release(struct inode *inode, struct file *file) 
{
	printk(KERN_INFO "+ seqfile: called release");
	return single_release(inode, file);
}

static struct proc_ops fops = {
	.proc_open = seqfile_open,
	.proc_read = seqfile_read,
	.proc_write = my_write,
	.proc_release = seqfile_release, 
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

static int __init init_seqfile_module(void) 
{	
	if (!(cookie_buffer = vmalloc(BUF_SIZE))) 
	{
		freemem();
		printk(KERN_ERR "+ seqfile: error during vmalloc");
		return -ENOMEM;
	}
	memset(cookie_buffer, 0, BUF_SIZE);
	if (!(fortune_dir = proc_mkdir(DIRNAME, NULL))) 
	{
		freemem();
		printk(KERN_ERR "+ seqfile: error during directory creation");
		return -ENOMEM;
	} 
	else if (!(fortune_file = proc_create(FILENAME, 0666, fortune_dir, &fops))) 
	{
		freemem();
		printk(KERN_ERR "+ seqfile: error during file creation");
		return -ENOMEM;
	} 
	else if (!(fortune_link = proc_symlink(SYMLINK, NULL, FILEPATH))) 
	{
		freemem();
		printk(KERN_ERR "+ seqfile: error during symlink creation");
		return -ENOMEM;
	}
	write_index = 0;
	read_index = 0;
	printk(KERN_INFO "+ module loaded");
	return 0;
}

static void __exit exit_seqfile_module(void) 
{
	freemem();
	printk(KERN_INFO "+ seqfile: unloaded");
}

module_init(init_seqfile_module);
module_exit(exit_seqfile_module);
