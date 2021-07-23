#include <stdio.h>
#include <string.h>

char *ciao(char *a,char b){
    int i;
    for ( i=0 ; i<strlen(a) ; i++){
        if (a[i] == b){
            return a+i;
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
