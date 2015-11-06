/* add_basic.c

   Demonstrates creating a DLL with an exported function, the inflexible way.
*/
//gcc -c -o add_basic.o add_basic.c
//gcc -o add_basic.dll -s -shared add_basic.o -Wl,--subsystem,windows
__declspec(dllexport) int __cdecl Add(int a, int b)
{
  return (a + b);
}