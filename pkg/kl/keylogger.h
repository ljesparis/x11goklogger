#ifndef X11GOKLOGGER_H
#define X11GOKLOGGER_H

extern unsigned int __init_kelogger       (const char* denv);
extern unsigned int __destroy_kelogger    ();
extern char*        __start_reading_input ();

#endif //X11GOKLOGGER_H