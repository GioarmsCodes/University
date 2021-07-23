#include <stdio.h>
#include <string.h>

char *mia(char *a,char *b){
    int i,j;
    for (i=0 ; i<strlen(a) ; i++){
        if (a[i] == b[0]){
            for (j=0 ; j<strlen(b) ; j++){
                if (a[j+i] != b[j]){
                    break;
                }
            }
            if (j==strlen(b)){
                return a+i;
            }
        }
    }
    return NULL;
}

int main(){

    char *main = "ciao come stai grandissimo";
    char *needle = "hello";

    printf("%s\n",strstr(main,needle));

    printf("%s\n",mia(main,needle));
}
