#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include"../library/lib.h"
int main() {
    float min=100,max=-100,mod;
    printf("╔════════════╦════════════╗\n");
    for (float k = 0; k < 11; k+= 0.11) {

        printf("║ %10f ║ %10f ║\n", k, function19(k));
        if(function19(k)>max){
            max=function19(k);
        }
        if(function19(k)<min){
            min=function19(k);
        }            
    }
    mod=fabs(max-min);
    printf("╠════════════╩════════════╣\n");
    printf("║ Module: %13f   ║\n", mod);
    printf("╚═════════════════════════╝\n");
    return 0;    
}    