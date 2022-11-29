#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include"../library/lib.h"
int main(){
    float y[8],negative[7],max=-100;
    int negative_num=0,max_index,index[7];
    for( int k=1;k<=7;k++){
        y[k]=function21(k);
        if(y[k]<0){
            negative[negative_num]=y[k];
            index[negative_num]=k;
            negative_num++;
        }
        if(y[k]>max){
          max=y[k];
          max_index=k;  
        }
    }
    for(int i=1;i<=7;i++){
        printf("%f\n",y[i]);
    }
    printf("\n");
   y[index[negative_num-2]]=y[max_index];
   for(int i=1;i<=7;i++){
        printf("%f\n",y[i]);
    }
    return 0;
   
}