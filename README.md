# Что?

gigfun - Утилита-костыль для видеокарты Gigabyte GeForce RTX 3090 Ti ([GV-N309TGAMING OC-24GD](https://www.gigabyte.com/Graphics-Card/GV-N309TGAMING-OC-24GD)).

# Зачем?

Видеокарта наотрез отказывается разгонять кулеры быстрее 80% даже когда начинает откровенно перегреваться, хотя переключатель BIOS стоит в положении "OC", а не "Silent". Данная утилита-костыль исправляет это недоразумение, форсируя запуск кулеров на 100% когда температура переваливает за 70℃.

# Как?

Демон должен стартовать вместе с X-ами поскольку требует опцию "Coolbits" (см. далее).

# Зависимости

- `nvidia-settings`;
- `nvidia-utils` (утилита `nvidia-smi`);

# Сборка

Требуется golang >= 1.16.

# Использование

1. Настраиваем опцию "Coolbits" как [написано в wiki](https://wiki.archlinux.org/title/NVIDIA/Tips_and_tricks#Enabling_overclocking):

   ```
   # /etc/X11/xorg.conf.d/nvidia.conf
   Section "Device"
       Identifier "Device0"
       Driver     "nvidia"
       Option     "Coolbits" "4"
   EndSection
   ```

2. Копируем юнит к текущему пользователю, обновляем список юнитов, запускаем юнит:

   ```
   install -D -m 0644 /usr/share/gigfun/service.example ~/.config/systemd/user/gigfun.service
   systemctl --user daemon-reload
   systemctl --user enable --now gigfun.service
   ```

# Лицензия

GPL.
