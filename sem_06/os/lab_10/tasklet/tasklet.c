#include <linux/kernel.h>
#include <linux/module.h>
#include <linux/interrupt.h>
#include <linux/slab.h>
#include <asm/io.h>

#include "ascii.h"

#define IRQ_NUMBER 1

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Glotov Ilya");

struct tasklet_struct *tasklet;
char tasklet_data[] = "key is pressed";

void tasklet_func(unsigned long data)
{
    printk(KERN_INFO "++: ***********************");
    printk(KERN_INFO "++: tasklet count = %u", tasklet->count.counter);
    printk(KERN_INFO "++: tasklet state = %lu", tasklet->state);

    printk(KERN_INFO "++: key code is %d", data);

    if (data < ASCII_LEN)
        printk(KERN_INFO "++: the key is %s", ascii[data]);

    printk(KERN_INFO "++: ***********************");
}

static irqreturn_t interrupt_handler(int irq, void *dev_id)
{
  int code;
  printk(KERN_INFO "++ interrupt_handler was called\n");
  if (irq == IRQ_NUMBER)
  {
    printk(KERN_INFO "++: tasklet state (before) = %lu",
                   tasklet->state);
    code = inb(0x60);
    tasklet->data = code;
    tasklet_schedule(tasklet);
    printk(KERN_INFO "++ tasklet was scheduled\n");
    printk(KERN_INFO "++: tasklet state (after) = %lu",
                    tasklet->state);
    return IRQ_HANDLED;
  }
  else
    return IRQ_NONE;
}

static int __init my_init(void) 
{
  tasklet = kmalloc(sizeof(struct tasklet_struct), GFP_KERNEL);

  if (tasklet == NULL)
  {
        printk(KERN_ERR "++: kmalloc error");
        return -1;
  }

  tasklet_init(tasklet, tasklet_func, (unsigned long)tasklet_data);

  if (request_irq(IRQ_NUMBER, interrupt_handler, IRQF_SHARED, "tasklet_interrupt_handler", interrupt_handler))
  {
    printk(KERN_ERR "++ Failed to register irq handler\n");
    return -1;
  }
  printk(KERN_INFO "++ Module loaded successfully\n");
  return 0;
}

static void __exit my_exit(void)
{
  tasklet_kill(tasklet);
  free_irq(IRQ_NUMBER, interrupt_handler);

  printk(KERN_INFO "++ Module unloaded\n");
}

module_init(my_init);
module_exit(my_exit);
