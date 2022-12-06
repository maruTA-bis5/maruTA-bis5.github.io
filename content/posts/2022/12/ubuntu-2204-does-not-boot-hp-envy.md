---
title: "HPのラップトップPCでUbuntu 22.04 (Jammy Jellyfish) が起動しない問題とその回避方法について"
date: 2022-12-06T23:45:00+0900
draft: false
description: "HPのラップトップPCでUbuntu 22.04 (Jammy Jellyfish) が起動しない問題とその回避方法について"
tags:
    - Ubuntu
    - jammy
    - HP ENVY
---

HPのラップトップPC (ENVY x360 Convertible 13-ar0074au)をWindows 11 (Insider Preview) + Ubuntu Desktopのデュアルブートにして、基本はUbuntuを使って生活しているのですが、重い腰を上げて20.04→22.04に更新したところ、高確率で起動中に固まるようになってしまいました。  
(20.04→22.04は`do-release-upgrade`を叩いただけで特筆することはない)

recovery modeで立ち上げてrootシェルに入ることは出来たので、ログを確認しつつGoogle先生に問い合わせたところ、以下のバグに該当することがわかりました。  
{{< linkCard "https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1977684" "Bug #1977684 “After Upgrade to 22.04. System does not boot into ...” : Bugs : linux package : Ubuntu" >}}  
{{< linkCard "https://lore.kernel.org/all/bug-215117-208809@https.bugzilla.kernel.org%2F/" "[Bug 215117] New: ucsi_acpi: kernel NULL pointer dereference" >}}

<details>

<summary>実際に`/var/kern.log`に出力されていたログ</summary>

```
Dec  6 22:12:51 capella kernel: [    3.369299] BUG: kernel NULL pointer dereference, address: 0000000000000058
Dec  6 22:12:51 capella kernel: [    3.370686] #PF: supervisor read access in kernel mode
Dec  6 22:12:51 capella kernel: [    3.371607] #PF: error_code(0x0000) - not-present page
Dec  6 22:12:51 capella kernel: [    3.372330] PGD 0 P4D 0 
Dec  6 22:12:51 capella kernel: [    3.373046] Oops: 0000 [#1] SMP NOPTI
Dec  6 22:12:51 capella kernel: [    3.373764] CPU: 0 PID: 6 Comm: kworker/0:0 Not tainted 5.15.0-56-generic #62-Ubuntu
Dec  6 22:12:51 capella kernel: [    3.374508] Hardware name: HP HP ENVY x360 Convertible 13-ar0xxx/85DE, BIOS F.26 08/29/2022
Dec  6 22:12:51 capella kernel: [    3.375267] Workqueue: events_long ucsi_init_work [typec_ucsi]
Dec  6 22:12:51 capella kernel: [    3.376052] RIP: 0010:typec_register_altmode+0x30/0x3b0 [typec]
Dec  6 22:12:51 capella kernel: [    3.376858] Code: 48 89 e5 41 57 41 56 41 55 49 89 f5 41 54 49 89 fc 48 8d bf 00 03 00 00 53 48 83 ec 30 65 48 8b 04 25 28 00 00 00 48 89 45 d0 <48> 8b 87 58 fd ff ff 48 3d 80 27 4c c0 74 1a 49 8d 94 24 f0 02 00
Dec  6 22:12:51 capella kernel: [    3.377727] RSP: 0018:ffffb0cb00107ca0 EFLAGS: 00010286
Dec  6 22:12:51 capella kernel: [    3.378603] RAX: 500dc2a12ea50000 RBX: ffff907ecbc61000 RCX: 0000000000000001
Dec  6 22:12:51 capella kernel: [    3.379498] RDX: 0000000000000000 RSI: ffffb0cb00107d78 RDI: 0000000000000300
Dec  6 22:12:51 capella kernel: [    3.380403] RBP: ffffb0cb00107cf8 R08: 0000000000000000 R09: 0000000000000000
Dec  6 22:12:51 capella kernel: [    3.381314] R10: 0000000000000003 R11: 0000000000000000 R12: 0000000000000000
Dec  6 22:12:51 capella kernel: [    3.382240] R13: ffffb0cb00107d78 R14: ffffb0cb00107d78 R15: 0000000000000000
Dec  6 22:12:51 capella kernel: [    3.383161] FS:  0000000000000000(0000) GS:ffff907f58a00000(0000) knlGS:0000000000000000
Dec  6 22:12:51 capella kernel: [    3.384103] CS:  0010 DS: 0000 ES: 0000 CR0: 0000000080050033
Dec  6 22:12:51 capella kernel: [    3.385047] CR2: 0000000000000058 CR3: 000000010b3aa000 CR4: 00000000003506f0
Dec  6 22:12:51 capella kernel: [    3.386013] Call Trace:
Dec  6 22:12:51 capella kernel: [    3.386976]  <TASK>
Dec  6 22:12:51 capella kernel: [    3.387954]  ? wait_for_completion_timeout+0x1d/0x30
Dec  6 22:12:51 capella kernel: [    3.388946]  typec_partner_register_altmode+0xe/0x20 [typec]
Dec  6 22:12:51 capella kernel: [    3.389957]  ucsi_register_altmode.constprop.0+0x1f0/0x280 [typec_ucsi]
Dec  6 22:12:51 capella kernel: [    3.390983]  ucsi_register_altmodes+0x156/0x210 [typec_ucsi]
Dec  6 22:12:51 capella kernel: [    3.392012]  ucsi_check_altmodes+0x1c/0x50 [typec_ucsi]
Dec  6 22:12:51 capella kernel: [    3.393030]  ucsi_register_port+0x4d4/0x510 [typec_ucsi]
Dec  6 22:12:51 capella kernel: [    3.394054]  ucsi_init+0xce/0x1b0 [typec_ucsi]
Dec  6 22:12:51 capella kernel: [    3.395089]  ucsi_init_work+0x16/0x30 [typec_ucsi]
Dec  6 22:12:51 capella kernel: [    3.396128]  process_one_work+0x22b/0x3d0
Dec  6 22:12:51 capella kernel: [    3.397178]  worker_thread+0x53/0x420
Dec  6 22:12:51 capella kernel: [    3.398242]  ? process_one_work+0x3d0/0x3d0
Dec  6 22:12:51 capella kernel: [    3.399289]  kthread+0x12a/0x150
Dec  6 22:12:51 capella kernel: [    3.400346]  ? set_kthread_struct+0x50/0x50
Dec  6 22:12:51 capella kernel: [    3.401409]  ret_from_fork+0x22/0x30
Dec  6 22:12:51 capella kernel: [    3.402476]  </TASK>
Dec  6 22:12:51 capella kernel: [    3.403541] Modules linked in: snd_hda_codec_generic ledtrig_audio amd64_edac(-) snd_hda_codec_hdmi edac_mce_amd amdgpu(+) snd_hda_intel snd_intel_dspcfg snd_intel_sdw_acpi snd_hda_codec rtw88_8822be kvm_amd rtw88_8822b snd_hda_core fjes(-) rtw88_pci kvm snd_hwdep crct10dif_pclmul snd_seq_midi nls_iso8859_1 rtw88_core snd_seq_midi_event ghash_clmulni_intel btusb iommu_v2 gpu_sched snd_rawmidi drm_ttm_helper aesni_intel ttm snd_pci_acp6x crypto_simd mac80211 btrtl cryptd btbcm snd_seq rapl input_leds btintel drm_kms_helper snd_pcm bluetooth snd_seq_device cec snd_timer rc_core i2c_algo_bit hp_wmi ecdh_generic fb_sys_fops serio_raw platform_profile wmi_bmof syscopyarea hid_multitouch(+) ecc k10temp snd_pci_acp5x cfg80211 sparse_keymap snd sysfillrect snd_rn_pci_acp3x sysimgblt ccp snd_pci_acp3x soundcore ucsi_acpi libarc4 typec_ucsi typec mac_hid hid_sensor_accel_3d hid_sensor_gyro_3d hid_sensor_magn_3d acpi_tad hid_sensor_trigger amd_pmc hp_accel industrialio_triggered_buffer
Dec  6 22:12:51 capella kernel: [    3.403601]  wireless_hotkey lis3lv02d kfifo_buf hid_sensor_iio_common industrialio sch_fq_codel overlay iptable_filter ip6table_filter ip6_tables br_netfilter bridge stp llc arp_tables msr parport_pc ppdev lp parport ramoops reed_solomon pstore_blk pstore_zone drm efi_pstore ip_tables x_tables autofs4 rtsx_pci_sdmmc nvme hid_sensor_hub xhci_pci hid_generic crc32_pclmul nvme_core i2c_piix4 amd_sfh rtsx_pci xhci_pci_renesas wmi i2c_hid_acpi i2c_hid video hid
Dec  6 22:12:51 capella kernel: [    3.412863] CR2: 0000000000000058
Dec  6 22:12:51 capella kernel: [    3.414342] ---[ end trace 041a0b479bb0ae03 ]---
Dec  6 22:12:51 capella kernel: [    3.494818] rtw_8822be 0000:02:00.0 wlo1: renamed from wlan0
Dec  6 22:12:51 capella kernel: [    3.698804] RIP: 0010:typec_register_altmode+0x30/0x3b0 [typec]
Dec  6 22:12:51 capella kernel: [    3.700304] Code: 48 89 e5 41 57 41 56 41 55 49 89 f5 41 54 49 89 fc 48 8d bf 00 03 00 00 53 48 83 ec 30 65 48 8b 04 25 28 00 00 00 48 89 45 d0 <48> 8b 87 58 fd ff ff 48 3d 80 27 4c c0 74 1a 49 8d 94 24 f0 02 00
Dec  6 22:12:51 capella kernel: [    3.701840] RSP: 0018:ffffb0cb00107ca0 EFLAGS: 00010286
Dec  6 22:12:51 capella kernel: [    3.703364] RAX: 500dc2a12ea50000 RBX: ffff907ecbc61000 RCX: 0000000000000001
Dec  6 22:12:51 capella kernel: [    3.704899] RDX: 0000000000000000 RSI: ffffb0cb00107d78 RDI: 0000000000000300
Dec  6 22:12:51 capella kernel: [    3.706399] RBP: ffffb0cb00107cf8 R08: 0000000000000000 R09: 0000000000000000
Dec  6 22:12:51 capella kernel: [    3.707893] R10: 0000000000000003 R11: 0000000000000000 R12: 0000000000000000
Dec  6 22:12:51 capella kernel: [    3.709360] R13: ffffb0cb00107d78 R14: ffffb0cb00107d78 R15: 0000000000000000
Dec  6 22:12:51 capella kernel: [    3.710808] FS:  0000000000000000(0000) GS:ffff907f58a00000(0000) knlGS:0000000000000000
Dec  6 22:12:51 capella kernel: [    3.712250] CS:  0010 DS: 0000 ES: 0000 CR0: 0000000080050033
Dec  6 22:12:51 capella kernel: [    3.713688] CR2: 0000000000000058 CR3: 000000010b3aa000 CR4: 00000000003506f0
```
</details>

### 発生する事象
一部のサービスの起動が完了せず、起動待ちの状態から進まなくなっていました。(カーネルパラメーターの`quiet`,`splash`を取り除いて起動するとわかりやすい)  
起動しないサービスは`systemd-hostnamed`, `systemd-oomd`, `systemd-logind`等ですが、毎回決まったサービスが起動しなくなるわけではなく増減するようです。

### 原因?
(Linuxカーネルのことはよくわからんのですが)`ucsi_acpi`モジュールにNULL参照バグがあるようです。

### 回避方法
`ucsi_acpi`モジュールをblacklistに入れて起動すれば(バグがあるモジュールを読み込まないので)起動できました。  
{{< linkCard "https://wiki.archlinux.jp/index.php/%E3%82%AB%E3%83%BC%E3%83%8D%E3%83%AB%E3%83%A2%E3%82%B8%E3%83%A5%E3%83%BC%E3%83%AB#.E3.83.96.E3.83.A9.E3.83.83.E3.82.AF.E3.83.AA.E3.82.B9.E3.83.88" "カーネルモジュール - ArchWiki" >}}

