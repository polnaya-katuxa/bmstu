#include <linux/module.h> 
#include <linux/kernel.h> 
#include <linux/init.h>  
#include <linux/vmalloc.h>
#include <linux/proc_fs.h> 
#include <linux/uaccess.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Karpova Ekaterina");

#define BUF_SIZE PAGE_SIZE

#define DIRNAME "fortunes"
#define FILENAME "fortune"
#define SYMLINK "fortune_link"
#define FILEPATH DIRNAME "/" FILENAME

static struct proc_dir_entry *fortune_dir = NULL;
static struct proc_dir_entry *fortune_file = NULL;
static struct proc_dir_entry *fortune_link = NULL;

static char *cookie_buffer;
static int write_index;
static int read_index;

static char tmp[BUF_SIZE];

ssize_t fortune_read(struct file *filp, char __user *buf, size_t count, loff_t *offp) 
{
	int len;
	printk(KERN_INFO "+ fortune: read called");
	if (*offp > 0 || !write_index) 
	{
		printk(KERN_INFO "+ fortune: empty");
		return 0;
	}
	if (read_index >= write_index)
		read_index = 0;
	len = snprintf(tmp, BUF_SIZE, "%s\n", &cookie_buffer[read_index]);
	if (copy_to_user(buf, tmp, len)) 
	{
		printk(KERN_ERR "+ fortune: copy_to_user error");
		return -EFAULT;
	}
	read_index += len;
	*offp += len;
	return len;
}

ssize_t fortune_write(struct file *filp, const char __user *buf, size_t len, loff_t *offp) 
{
	printk(KERN_INFO "+ fortune: write called");
	if (len > BUF_SIZE - write_index + 1) 
	{
		printk(KERN_ERR "+ fortune: cookie_buffer overflow");
		return -ENOSPC;
	}
	if (copy_from_user(&cookie_buffer[write_index], buf, len)) 
	{
		printk(KERN_ERR "+ fortune: copy_to_user error");
		return -EFAULT;
	}
	write_index += len;
	cookie_buffer[write_index - 1] = '\0';
	return len;
}

int fortune_open(struct inode *inode, struct file *file) 
{
	printk(KERN_INFO "+ fortune: called open");
	return 0;
}

int fortune_release(struct inode *inode, struct file *file) 
{
	printk(KERN_INFO "+ fortune: called release");
	return 0;
}

static const struct proc_ops fops = {
	proc_read: fortune_read, 
	proc_write: fortune_write, 
	proc_open: fortune_open, 
	proc_release: fortune_release
};

static void freemem(void) 
{
	if (fortune_link)
		remove_proc_entry(SYMLINK, NULL);
	if (fortune_file)
		remove_proc_entry(FILENAME, fortune_dir);
	if (fortune_dir)
		remove_proc_entry(DIRNAME, NULL);
	if (cookie_buffer)
		vfree(cookie_buffer);
}

static int __init fortune_init(void) 
{
	if (!(cookie_buffer = vmalloc(BUF_SIZE))) 
	{
		freemem();
		printk(KERN_ERR "+ fortune: error during vmalloc");
		return -ENOMEM;
	}
	memset(cookie_buffer, 0, BUF_SIZE);
	if (!(fortune_dir = proc_mkdir(DIRNAME, NULL))) 
	{
		freemem();
		printk(KERN_ERR "+ fortune: error during directory creation");
		return -ENOMEM;
	} 
	else if (!(fortune_file = proc_create(FILENAME, 0666, fortune_dir, &fops))) 
	{
		freemem();
		printk(KERN_ERR "+ fortune: error during file creation");
		return -ENOMEM;
	} 
	else if (!(fortune_link = proc_symlink(SYMLINK, NULL, FILEPATH))) 
	{
		freemem();
		printk(KERN_ERR "+ fortune: error during symlink creation");
		return -ENOMEM;
	}
	write_index = 0;
	read_index = 0;
	printk(KERN_INFO "+ fortune: module loaded");
	return 0;
}

static void __exit fortune_exit(void) 
{
	freemem();
	printk(KERN_INFO "+ fortune: module unloaded");
}

module_init(fortune_init) 
module_exit(fortune_exit)
