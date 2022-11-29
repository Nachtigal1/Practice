#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include"../library/lib.h"
int main() {
    float y, business[2] = {0, 0};
    for (int k = 1991; k <= 2001; k++) {
        y = 100 * function20(k);
        printf("%f\n", y);

        if ((y >230)&&(y < 8500)) {
            business[0]++;
            business[1]+=y;       
        }
    }
    printf("Years with+: %f\n All years with +: %f\n", business[1],  business[0]);
    return 0;
}