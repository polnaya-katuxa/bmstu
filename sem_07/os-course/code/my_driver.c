#include <linux/module.h>
#include <linux/init.h>
#include <linux/fs.h>
#include <asm/io.h>
#include <linux/unistd.h>
#include <linux/delay.h>
#include <asm/uaccess.h>
#include <linux/miscdevice.h>
#include <linux/fcntl.h>
#include <linux/ioctl.h>
#include <linux/i2c-dev.h>
#include <linux/types.h>
#include <linux/delay.h>
#include <linux/i2c.h>
#include <uapi/linux/errno.h>
#include <asm/errno.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Karpova Ekaterina");

#define LCD_ADDRESS     (0x3e)
#define RGB_ADDRESS     (0xc0>>1)
#define COMMAND_REG 0x80
#define DATA_REG 0x40

#define REG_RED     0x04
#define REG_GREEN       0x03
#define REG_BLUE        0x02
#define REG_MODE1       0x00
#define REG_MODE2       0x01
#define REG_OUTPUT      0x08
#define LCD_CLEARDISPLAY  0x01
#define LCD_ENTRYMODESET  0x04
#define LCD_DISPLAYCONTROL  0x08
#define LCD_FUNCTIONSET  0x20

//#flags for display entry mode
#define LCD_ENTRYLEFT  0x02
#define LCD_ENTRYSHIFTDECREMENT  0x00

//#flags for display on/off control
#define LCD_DISPLAYON  0x04
#define LCD_CURSOROFF  0x00

//#flags for function set
#define LCD_4BITMODE  0x00
#define LCD_2LINE  0x08

struct file* file_i2c;

static inline int delay_microseconds(int value) {
    if (value > 1000) {
        msleep(value/1000);
    }
    udelay(value % 1000);
    return 0;
}

int write_byte_data(int reg, int data, const int addr) {
    if (vfs_ioctl(file_i2c, I2C_SLAVE, addr) < 0)
    {
        printk(KERN_DEBUG "+ failed to acquire bus access and/or talk to slave");
        return 1;
    }

    union i2c_smbus_data msg;
    msg.byte = data;

    struct i2c_smbus_ioctl_data req = {
        .read_write = 0,
        .command = reg,
        .size = 2,
        .data = &msg,
    };

    struct i2c_client *client = file_i2c->private_data;
    int res = i2c_smbus_xfer(client->adapter, client->addr, client->flags,
                         req.read_write, req.command, req.size, &msg);
    if (res < 0) {
        printk(KERN_DEBUG "+ failed to smbus_xfer %d", res);
        return 1;
    }

    return 0;
}

int command(int cmd) {
    if ((write_byte_data(COMMAND_REG, cmd, LCD_ADDRESS)) != 0)
    {
        printk(KERN_DEBUG "+ failed to write byte data");
        return 1;
    }

    return 0;
}

int print_byte(int data) {
    if ((write_byte_data(DATA_REG, data, LCD_ADDRESS)) != 0)
    {
        printk(KERN_DEBUG "+ failed to write byte data");
        return 1;
    }

    return 0;
}

int set_reg(int reg, int data) {
    if ((write_byte_data(reg, data, RGB_ADDRESS)) != 0)
    {
        printk(KERN_DEBUG "+ failed to write byte data");
        return 1;
    }

    return 0;
}

void set_cursor(int col, int row) {
    if (row == 0) {
        col |= 0x80;
    } else {
        col |= 0xc0;
    }

    if ((command(col)) != 0) {
        printk(KERN_DEBUG "+ failed to set cursor");
    }
}

int configure(void) {
    printk(KERN_DEBUG "+ set show_function 1st try");

    if ((command(LCD_4BITMODE | LCD_2LINE | LCD_FUNCTIONSET)) != 0) {
        printk(KERN_DEBUG "+ failed to set show_function 1st try");
        return 1;
    }

    delay_microseconds(100000);

    printk(KERN_DEBUG "+ set show_function 2nd try");

    if ((command(LCD_4BITMODE | LCD_2LINE | LCD_FUNCTIONSET)) != 0) {
        printk(KERN_DEBUG "+ failed to set show_function 2nd try");
        return 1;
    }

    delay_microseconds(100000);

    printk(KERN_DEBUG "+ set show_function 3rd try");

    if ((command(LCD_4BITMODE | LCD_2LINE | LCD_FUNCTIONSET)) != 0) {
        printk(KERN_DEBUG "+ failed to set show_function 3rd try");
        return 1;
    }

    printk(KERN_DEBUG "+ set show_function 4th try");

    if ((command(LCD_4BITMODE | LCD_2LINE | LCD_FUNCTIONSET)) != 0) {
        printk(KERN_DEBUG "+ failed to set show_function 4th try");
        return 1;
    }

    delay_microseconds(100000);

    printk(KERN_DEBUG "+ turn the display on with no cursor or blinking default");

    if ((command(LCD_DISPLAYON | LCD_CURSOROFF | LCD_DISPLAYCONTROL)) != 0) {
        printk(KERN_DEBUG "+ failed to turn the display on with no cursor or blinking default");
        return 1;
    }

    printk("+ write to send command clear");

    if ((command(LCD_CLEARDISPLAY)) != 0) {
        printk(KERN_DEBUG "+ failed to write to send command clear");
        return 1;
    }

    delay_microseconds(100000);

    printk(KERN_DEBUG "+ set text direction and entry mode");

    if ((command(LCD_ENTRYLEFT | LCD_ENTRYSHIFTDECREMENT | LCD_ENTRYMODESET)) != 0) {
        printk(KERN_DEBUG "+ failed to set text direction and entry mode");
        return 1;
    }

    printk(KERN_DEBUG "+ setting colors");

    set_reg(REG_MODE1, 0);
    set_reg(REG_OUTPUT, 0xFF);
    set_reg(REG_MODE2, 0x20);

    set_reg(REG_RED, 255);
    set_reg(REG_GREEN, 255);
    set_reg(REG_BLUE, 255);

    printk(KERN_DEBUG "+ ending configuration");

    return 0;
}

int clear(void) {
    if ((command(LCD_CLEARDISPLAY)) != 0) {
        printk(KERN_DEBUG "+ failed to write to send command");
        return 1;
    }

    return 0;
}

int print_string(char* str) {
    for (int i = 0; str[i] != 0; i++) {
        if (i > 0 && i % 16 == 0) {
            set_cursor(0, 1);
        }
        if ((print_byte(str[i])) != 0) {
            printk(KERN_DEBUG "+ Failed to write data");
            return 1;
        }
    }

    return 0;
}

int start(void) {
    printk(KERN_DEBUG "+ opening");

    char *filename = (char*)"/dev/i2c-5";
    file_i2c = filp_open(filename, O_RDWR, 0);
    if (!file_i2c) {
        printk(KERN_DEBUG "+ failed to open the i2c bus");
        return 1;
    }

    printk(KERN_DEBUG "+ configuring");

    if ((configure()) != 0) {
        printk(KERN_DEBUG "+ failed to configure");
        return 1;
    }

    printk(KERN_DEBUG "+ setting cursor");

    set_cursor(0, 0);

    delay_microseconds(100000);

    return 0;
}

static char *info_str = "hello, world\n";

static ssize_t dev_read( struct file * file, char * buf, size_t count, loff_t *ppos) {
    int len = strlen(info_str);

    if (count < len) {
        return -EINVAL;
    }

    if (*ppos != 0) {
        return 0;
    }

    if (copy_to_user(buf, info_str, len)) {
        return -EINVAL;
    }

    *ppos = len;
    return len;
}

static int str_pos = 0;
static int col_pos = 0;

static ssize_t dev_write(struct file *file, const char *buf, size_t count, loff_t *ppos) {
    char data[1024];
    if (copy_from_user(data, buf, strlen(buf))) {
        return -EINVAL;
    }

    str_pos = 0;
    col_pos = 0;

    for (int i = 0; i < count; i++) {
        if ((col_pos==0) && (str_pos==0)) {
            clear();
        }

        set_cursor(col_pos, str_pos);
        delay_microseconds(10000);

        printk(KERN_DEBUG "+ cursor col %d row %d cur sym %c", col_pos, str_pos, data[i]);

        if (buf[i] != '\n') {
            print_byte(data[i]);
            col_pos++;
            delay_microseconds(10000);
        } else {
            col_pos=16;
        }

        if (col_pos == 16) {
            col_pos = 0;
            str_pos++;
            if (str_pos == 2) {
                str_pos = 0;
            }
        }
    }

    return count;
}

static const struct file_operations my_fops = {
        .owner  = THIS_MODULE,
        .read   = dev_read,
        .write  = dev_write,
};

static struct miscdevice my_dev = {
        MISC_DYNAMIC_MINOR,
        "proclcd",
        &my_fops
};

static int __init my_init(void)
{
    int rc = misc_register(&my_dev);
    if (rc) {
        printk(KERN_ERR "+ unable to register misc device\n");
    }

    start();

    printk(KERN_DEBUG "+ module loaded");

    return rc;
}

static void __exit my_exit(void)
{
    misc_deregister( &my_dev );
    printk(KERN_DEBUG "+ module unloaded");
}

module_init(my_init);
module_exit(my_exit);
