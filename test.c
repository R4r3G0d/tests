#include <stdio.h>
#include <string.h>

void vulnerableFunction(const char *input) {
    char buffer[100];
    strcpy(buffer, input); 
}

int main() {
    char userInput[200];
    fgets(userInput, sizeof(userInput), stdin);
    vulnerableFunction(userInput);
    return 0;
}
