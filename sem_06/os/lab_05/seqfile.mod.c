#include <linux/module.h>
#define INCLUDE_VERMAGIC
#include <linux/build-salt.h>
#include <linux/elfnote-lto.h>
#include <linux/vermagic.h>
#include <linux/compiler.h>

BUILD_SALT;
BUILD_LTO_INFO;

MODULE_INFO(vermagic, VERMAGIC_STRING);
MODULE_INFO(name, KBUILD_MODNAME);

__visible struct module __this_module
__section(".gnu.linkonce.this_module") = {
	.name = KBUILD_MODNAME,
	.init = init_module,
#ifdef CONFIG_MODULE_UNLOAD
	.exit = cleanup_module,
#endif
	.arch = MODULE_ARCH_INIT,
};

#ifdef CONFIG_RETPOLINE
MODULE_INFO(retpoline, "Y");
#endif

static const struct modversion_info ____versions[]
__used __section("__versions") = {
	{ 0x76a006e9, "module_layout" },
	{ 0x46fa162b, "proc_symlink" },
	{ 0x3acce11f, "proc_create" },
	{ 0xcd1bfa5, "proc_mkdir" },
	{ 0xd6ee688f, "vmalloc" },
	{ 0x6863141c, "remove_proc_entry" },
	{ 0x999e8297, "vfree" },
	{ 0x50c87208, "single_release" },
	{ 0x3195d15d, "seq_printf" },
	{ 0x656e4a6e, "snprintf" },
	{ 0x7696c2a6, "single_open" },
	{ 0x92a11473, "seq_read" },
	{ 0xdcb764ad, "memset" },
	{ 0x12a4e128, "__arch_copy_from_user" },
	{ 0x88db9f48, "__check_object_size" },
	{ 0x92997ed8, "_printk" },
	{ 0xb788fb30, "gic_pmr_sync" },
	{ 0x4b0a3f52, "gic_nonsecure_priorities" },
	{ 0x7b4627a9, "cpu_hwcap_keys" },
	{ 0x622f73ad, "cpu_hwcaps" },
	{ 0x14b89635, "arm64_const_caps_ready" },
};

MODULE_INFO(depends, "");


MODULE_INFO(srcversion, "D496673C6E8B7506D0A84B7");
