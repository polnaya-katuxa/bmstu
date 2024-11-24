cmd_/home/muhomorfus/OS/lab_06/modules.order := {   echo /home/muhomorfus/OS/lab_06/kittyfs.ko; :; } | awk '!x[$$0]++' - > /home/muhomorfus/OS/lab_06/modules.order
