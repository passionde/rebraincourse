# Задание GO-02 07

1. Создайте новую директорию с файлом main.go.
2. Скачайте файл in.txt и скопируйте его в директорию data рядом с файлом main.go.
3. Напишите программу, которая считывает данные из файла, проверяет, что все поля заполнены и записывает считанные данные в файл data/data_out.txt в
   формате Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n.
4. Если какое-то поле не заполнено, то программа должна вызвать панику, передав строку вида parse error: empty field on string %d.
5. Объявите необходимые отложенные вызовы.
6. Обработайте панику таким образом, чтобы после обработки на экран вывелось содержимое файла data_out.txt, которое было записано до возникновения паники.