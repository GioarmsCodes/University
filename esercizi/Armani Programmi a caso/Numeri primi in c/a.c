#include <stdio.h>

int Isprimo(int n){
    int i;
    for (i = 2 ;i<n ; i++){
        if (n%i == 0){
            return 0;
        }
    }
    return 1;
}

int main(){

    int x[1000];
    int n,i,c;
    c = 0;
    printf("inserisci numero: ");
    scanf("%d",&n);
    for (i = 2 ; i<= n ; i++){
        if (Isprimo(i) == 1){
            x[c] = i;
            c++;
        }

    }

    for (i = 0 ; i< c ; i++){
    printf("%d\n",x[i]);
    }
    printf("i numeri primi fino a %d sono : %d",n,c);
    return 0;
}
