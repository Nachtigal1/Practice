#include <stdio.h>

int main() {
    int rows;
    float value, valueChange, cheldron, liter, pitch;

    scanf("%i %f %f", &rows, &value, &valueChange);
    printf("╔══════════════════════════╦════════════════════════╦══════════════════════════╗\n");
    for (int i = 0; i < rows; i++, value += valueChange) {
        cheldron = value, liter= value * 1.309, pitch = value * 0.149;

        printf("║%20f  %19f %19f ║\n", cheldron, liter, pitch);
    }
    printf("╚══════════════════════════╩════════════════════════╩══════════════════════════╝\n");

    return 0;
}