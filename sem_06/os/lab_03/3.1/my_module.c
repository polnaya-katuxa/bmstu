#include <linux/init.h>
#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/sched.h>
#include <linux/init_task.h>
#include <linux/fs_struct.h>
#include <linux/path.h>

MODULE_AUTHOR("Katerina Karpova");
MODULE_DESCRIPTION("LAB");
MODULE_LICENSE("GPL");

void print_task(struct task_struct *task)
{
	printk(KERN_INFO "- %s, pid - %d, parent - %s, ppid - %d, state - %d, flags - %d, prio - %d, policy - %d, mnt - %s\n",
		   task->comm,
		   task->pid, 
		   task->parent->comm, 
		   task->parent->pid,
		   task->__state,
		   task->flags,
		   task->prio,
		   task->policy,
		   task->fs->root.dentry->d_iname);
}

static int __init m_init(void)
{
	struct task_struct *task = &init_task;
	do
	{
		print_task(task);
	} while ((task = next_task(task)) != &init_task);

	print_task(current);

	return 0;
}

static void __exit m_exit(void) 
{}

module_init(m_init);
module_exit(m_exit);


