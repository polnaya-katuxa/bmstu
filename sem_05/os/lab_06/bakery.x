/*
 * filename: bakery.x
     * function: Define constants, non-standard data types and the calling process in remote calls
 */

const GET_NUM = 0;
const OPEN_CRIT_SECTION = 1;

struct BAKERY
{
    int op;
    int id;
    int pid;
    int num;
    int res;
};

program BAKERY_PROG
{
    version BAKERY_VER
    {
        struct BAKERY BAKERY_PROC(struct BAKERY) = 1;
    } = 1; /* Version number = 1 */
} = 0x20000001; /* RPC program number */