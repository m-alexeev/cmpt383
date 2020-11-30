#include <stdio.h>




void printList(const char** list, int length){
    for (int i = 0; i < length; i ++){
        printf("%s\n", *(list+i));
    }
}