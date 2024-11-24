#include <linux/kernel.h>
#include <linux/vmalloc.h>
#include <linux/module.h>
#include <linux/errno.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Khudyakov Vladimir");

#define BUF_SIZE PAGE_SIZE

#define DIRNAME "fortune_seq_dir"
#define SUBDIRNAME "fortune_seq_subdir"
#define FILENAME "fortune_seq"
#define SYMLINK "fortune_seq_ln"
#define FILEPATH DIRNAME "/" SUBDIRNAME "/" FILENAME

static struct proc_dir_entry *dir = NULL;
static struct proc_dir_entry *subdir = NULL;
static struct proc_dir_entry *afile = NULL;
static struct proc_dir_entry *link = NULL;

static char *cookie_pot;
static int cookie_index;
static int next_fortune;

static char tmp[BUF_SIZE];

int seqfile_show(struct seq_file *m, void *v)
{
	int len;
	if (!cookie_index)
		return 0;
	if (next_fortune >= cookie_index)
		next_fortune = 0;
	len = snprintf(tmp, BUF_SIZE, "%s", cookie_pot + next_fortune);
	seq_printf(m, "%s", tmp);
	next_fortune += len + 1;

	printk("+ show() was called\n");

	return 0;
}

ssize_t seqfile_write(struct file *file, const char __user *buf, size_t len, loff_t *offp)
{
	printk("+ write() called\n");

	if (len > BUF_SIZE - cookie_index + 1)
	{
		printk(KERN_ERR"+ buffer overflow!\n");
		return -ENOSPC;
	}

	if (copy_from_user(cookie_pot + cookie_index, buf, len))
	{
		printk(KERN_ERR"+ copy_from_user() failed\n");
		return -EFAULT;
	}

	cookie_index += len;
	cookie_pot[cookie_index++] = 0;

	return len;
}

int seqfile_open(struct inode *inode, struct file *file)
{
	printk(KERN_INFO"+ open() called\n");
	return single_open(file, seqfile_show, NULL);
}

int seqfile_release(struct inode *inode, struct file *file)
{
	printk(KERN_INFO"+ release() called\n");
	return single_release(inode, file);
}

ssize_t seqfile_read(struct file *file, char __user *buf, size_t size, loff_t *ppos)
{	
	printk("+ read() called\n");
	return seq_read(file, buf, size, ppos);
}

static struct proc_ops fops = {
	.proc_read = seqfile_read,
	.proc_write = seqfile_write,
	.proc_open = seqfile_open,
	.proc_release = seqfile_release
};

static void freemem(void)
{
	if (link)
		remove_proc_entry(SYMLINK, dir);
	if (afile)
		remove_proc_entry(FILENAME, subdir);
	if (subdir)
		remove_proc_entry(SUBDIRNAME, dir);
	if (dir)
		remove_proc_entry(DIRNAME, NULL);
	if (cookie_pot)
		vfree(cookie_pot);
}

static int __init mod_init(void)
{
	if (!(cookie_pot = vmalloc(BUF_SIZE)))
	{
		printk(KERN_ERR"+ malloc failed!\n");
		return -1;
	}
	
	memset(cookie_pot, 0, BUF_SIZE);
	if (!(dir = proc_mkdir(DIRNAME, NULL)))
	{
		freemem();
		printk(KERN_ERR"+ mkdir failed!\n");
		return -1;
	}
	if (!(subdir = proc_mkdir(SUBDIRNAME, dir)))
	{
		freemem();
		printk(KERN_ERR"+ mkdir failed!\n");
		return -1;
	}
	if (!(afile = proc_create(FILENAME, 0666, subdir, &fops)))
	{
		freemem();
		printk(KERN_ERR"+ file create failed!\n");
		return - 1;
	}
	if (!(link = proc_symlink(SYMLINK, dir, FILEPATH)))
	{
		freemem();
		printk(KERN_ERR"+ failed to create symlink!\n");
		return -1;
	}

	cookie_index = 0;
	next_fortune = 0;
	printk(KERN_INFO"+ module loaded!\n");
	return 0;
}

static void __exit mod_exit(void)
{
	freemem();
	printk(KERN_INFO"+ module unloaded!\n");
}

module_init(mod_init);
module_exit(mod_exit);
