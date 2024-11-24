# Проект Тераграф. Тестовый пример проекта RiscV в составе Graph Rocessor Core

## Общее описание

Тестовый проект для взаимодействия HOST систем с Graph Processor Core (XRT Runtime версия), используются аппаратные FIFO буферы по 512 слов и память External Memory.

## Установка

Для установки требуется рекурсивно клонировать репозиторий:

```bash
git clone --recursive https://gitlab.com/leonhard-x64-xrt-v2/disc-example.git
```

## Зависимости

Зависимости для сборки проекта:

* набор средст сборки [riscv toolchain](https://gitlab.com/quantr/toolchain/riscv-gnu-toolchain) и экспорт исполняемых файлов в `PATH`

* набор библиотек [picolib](https://github.com/picolibc/picolibc) и экспорт в `C_INCLUDE_PATH`

* исходный текст проекта [taiga](https://github.gitop.top/taiga-project/taiga) и экспорт в переменную окружения `TAIGA_DIR`

* библиотека [xrt](https://gitlab.com/xilinx4jet/XRT) и установка по пути `/opt/xilinx/xrt`

Для стандартного пользователя ВМ студенческой команды хакатона все необходимые переменные окружения установлены по-умолчанию.

## Сборка проекта

Следует выполнить команду:

```bash
make
```

Результатом выполнения команды станет файлы host_main, sw_kernel_main.rawbinary и leonhard_2cores_267mhz.xclbin в директории проекта верхнего уровня.

| :exclamation:  Не забывайте синхронизировать тексты исходного текста host и kernel составляющих проекта |
|---------------------------------------------------------------------------------------------------------|

## Запуск проекта

```
./host_main leonhard_2cores_267mhz.xclbin sw_kernel_main.rawbinary
```

## Очистка проекта

Следует выполнить команду:

```bash
make clean
```
