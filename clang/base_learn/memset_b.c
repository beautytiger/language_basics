// C program to demonstrate working of memset() 
// https://www.geeksforgeeks.org/memset-c-example/
#include <stdio.h> 
#include <string.h> 
  
void printArray(int arr[], int n) 
{ 
   for (int i=0; i<n; i++) 
      printf("%d ", arr[i]); 
} 
  
int main() 
{ 
    int n = 10; 
    int arr[n]; 
  
    // Fill whole array with 0. 
    memset(arr, 0, n*sizeof(arr[0])); 
    printf("Array after memset()\n"); 
    printArray(arr, n); 
  
    return 0; 
}
