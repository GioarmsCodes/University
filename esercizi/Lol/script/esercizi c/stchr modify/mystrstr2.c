#include <stdio.h>
#include <string.h>

char *ciao(char *a,char b){
    int i;
    for ( ; *a!='\0' ; a++){
        if (*a == b){
            return a;
        }
    }
    return NULL;
}

int main(){

    char *word = "ciao";

    char find = 'i';

    printf("%s\n",strchr(word,find));
    printf("%s\n",ciao(word,find));


    return 0;
}
