#include <linux/build-salt.h>
#include <linux/module.h>
#include <linux/vermagic.h>
#include <linux/compiler.h>

BUILD_SALT;

MODULE_INFO(vermagic, VERMAGIC_STRING);
MODULE_INFO(name, KBUILD_MODNAME);

__visible struct module __this_module
__attribute__((section(".gnu.linkonce.this_module"))) = {
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
__used
__attribute__((section("__versions"))) = {
	{ 0x14da35b1, "module_layout" },
	{ 0x542a9844, "single_release" },
	{ 0xd97d8cf1, "seq_lseek" },
	{ 0x2bb6e5fe, "remove_proc_entry" },
	{ 0x62bcbcad, "proc_create_data" },
	{ 0xd6ee688f, "vmalloc" },
	{ 0x7c32d0f0, "printk" },
	{ 0x96acfa69, "single_open" },
	{ 0xa0ee1609, "PDE_DATA" },
	{ 0xbdfb6dbb, "__fentry__" },
};

static const char __module_depends[]
__used
__attribute__((section(".modinfo"))) =
"depends=";


MODULE_INFO(srcversion, "3499AD071378A48C7E08AD2");