#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/init_task.h>
#include <linux/fs.h>
#include <linux/slab.h>
#include <linux/time.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Aleksey Knyazhev");

#define MAGIC_NUM 0x12528391

#define CACHE_SIZE 1024
#define CACHE_NAME "kittyfs_cache"
static struct kmem_cache *cache = NULL;
static struct kittyfs_inode **inode_cache = NULL;
static size_t cache_index = 0;

static struct kittyfs_inode
{
    int i_mode;
    unsigned long i_ino;
} kittyfs_inode;

static void kittyfs_kill_sb(struct super_block *sb){
    printk(KERN_INFO "+ kittyfs: kill super block");
    kill_anon_super(sb);
}

static void kittyfs_put_sb(struct super_block *sb)
{
    printk(KERN_INFO "+ kittyfs: superblock destroy called");
}

static struct super_operations const kittyfs_sb_ops = {
    .put_super = kittyfs_put_sb,
    .statfs = simple_statfs,
    .drop_inode = generic_delete_inode,
};

static struct inode *kittyfs_new_inode(struct super_block *sb, int ino, int mode)
{
	struct inode *res;
	res = new_inode(sb);
	if (!res)
		return NULL;
	res->i_ino = ino;
    res->i_mode = mode;
    res->i_atime = res->i_mtime = res->i_ctime = current_time(res);
    res->i_op = &simple_dir_inode_operations;
    res->i_fop = &simple_dir_operations;
    inode_cache[cache_index] = kmem_cache_alloc(cache, GFP_KERNEL);
    if (inode_cache[cache_index])
    {
    	inode_cache[cache_index]->i_ino = res->i_ino;
    	inode_cache[cache_index]->i_mode = res->i_mode;
    	cache_index++;
    }
    return res;
}

static int kittyfs_fill_sb(struct super_block *sb, void *data, int silent)
{
    struct dentry *root_dentry;
    struct inode *root_inode;
    sb->s_blocksize = PAGE_SIZE;       
    sb->s_blocksize_bits = PAGE_SHIFT;
    sb->s_magic = MAGIC_NUM;
    sb->s_op = &kittyfs_sb_ops;
    root_inode = kittyfs_new_inode(sb, 1, S_IFDIR | 0755);
    if (!root_inode)
    {
        printk(KERN_INFO "+ kittyfs: cannot make root inode");
        return -ENOMEM;
    }
    root_dentry = d_make_root(root_inode);
    if (!root_dentry)
    {
        printk(KERN_ERR "+ kittyfs: cannot make root dentry");
        return -ENOMEM;
    }
    sb->s_root = root_dentry;
    return 0;
}

static struct dentry *kittyfs_mount(struct file_system_type *type, int flags, const char *dev, void *data)
{
    struct dentry *const root_dentry = mount_nodev(type, flags, data, kittyfs_fill_sb);
    if (IS_ERR(root_dentry))
        printk(KERN_ERR "+ kittyfs: cannot mount");
    else
        printk(KERN_INFO "+ kittyfs: mount successful");
    return root_dentry;
}

static void kittyfs_slab_constructor(void *addr)
{
    memset(addr, 0, sizeof(struct kittyfs_inode));
}

static struct file_system_type kittyfs_type = {
    .owner = THIS_MODULE,
    .name = "kittyfs",
    .mount = kittyfs_mount,
    .kill_sb = kittyfs_kill_sb,
};

static int __init kittyfs_init(void)
{
    int err = register_filesystem(&kittyfs_type);

    if (err != 0){
        printk(KERN_ERR "+ kittyfs: cannot register filesystem");
        return err;
    }

    if ((inode_cache = kmalloc(sizeof(struct kittyfs_inode*)*CACHE_SIZE, GFP_KERNEL)) == NULL)
    {
        printk(KERN_ERR "+ kittyfs: Can't kmalloc.\n");
        return -ENOMEM;
    }

    if ((cache = kmem_cache_create(CACHE_NAME, sizeof(struct kittyfs_inode), 0, SLAB_HWCACHE_ALIGN, kittyfs_slab_constructor)) == NULL)
    {
        printk(KERN_ERR "+ kittyfs: cannot create cache");
        kmem_cache_destroy(cache);
        kfree(inode_cache); 
        return -ENOMEM;
    }

    printk(KERN_INFO "+ kittyfs: module loaded");
    return 0;
}

static void __exit kittyfs_exit(void)
{
    int err;
    int i;
    
    for (i = 0; i < cache_index; i++)
    	kmem_cache_free(cache, inode_cache[i]);
    
    kmem_cache_destroy(cache);
    kfree(inode_cache);

    err = unregister_filesystem(&kittyfs_type);
    if (err != 0)
        printk(KERN_ERR "+ kittyfs: cannot unregister filesystem");
    else
        printk(KERN_INFO "+ kittyfs: module is unloaded");
}

module_init(kittyfs_init);
module_exit(kittyfs_exit);


