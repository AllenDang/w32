/* addtest_basic.c

   Demonstrates using the function imported from the DLL, the inelegant way.
*/
//gcc -c -o addtest_basic.o addtest_basic.c
//gcc -o addtest_basic.exe -s addtest_basic.o -L. -ladd_basic
   
#include <stdlib.h>
#include <stdio.h>

/* Declare imported function so that we can actually use it. */
__declspec(dllimport) int __cdecl Add(int a, int b);

int main(int argc, char** argv)
{
  printf("%d\n", Add(6, 23));

  return EXIT_SUCCESS;
}