#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Complete the following function.

int marks_summation(int *marks, int number_of_students, char gender) {
    // Write your code here.
    int start = gender == 'b' ? 0 : 1;
    int sum = 0;
    for (int i = start; i < number_of_students; i += 2) {
        sum += marks[i];
    }
    return sum;
}

int main(int argc, char *argv[]) {
    int number_of_students;
    char gender;
    int sum;

    scanf("%d", &number_of_students);
    int *marks = (int *)malloc(number_of_students * sizeof(int));

    for (int student = 0; student < number_of_students; student++) {
        scanf("%d", (marks + student));
    }

    scanf(" %c", &gender);
    sum = marks_summation(marks, number_of_students, gender);
    printf("%d", sum);
    free(marks);

    return 0;
}
