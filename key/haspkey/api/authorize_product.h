#ifndef _AUTHORIZE_PRODUCT_HEADER_INCLUDED
#define _AUTHORIZE_PRODUCT_HEADER_INCLUDED


#ifdef AUTHORIZE_EXPORTS
#define AUTHORIZE_API __declspec(dllexport)
#else
#define AUTHORIZE_API
#endif // AUTHORIZE_EXPORTS


// UKey不支持该产品
#define PRODUCT_NOKEY static_cast<unsigned>(-1)


#if defined (__cplusplus)
extern "C" {
#endif


///@brief
///判断Ukey是否支持该产品
///@param
///ucProduct-产品编号，1-255
///@return
///1: 支持该产品， 0: 不支持该产品
AUTHORIZE_API int cwIsProduct(unsigned char ucProduct);


///@brief
///获取产品的参数值
///@param
///ucProduct-产品编号，1-255
///@return
///PRODUCT_NOKEY: Ukey不支持该产品， 其他值: 该产品的参数值
AUTHORIZE_API unsigned cwGetProductParameter(unsigned char ucProduct);


///@brief
///将产品参数分解为字节数组
///@param
///ucProduct-产品编号，1-255
///ucByte-返回的字节数组，需要分配4个空间，字节从高位到地位依次为0，1，2，3
///@return
///0: 分解完成， -1: Ukey不支持该产品， -2: ucByte为空，没有分配空间
AUTHORIZE_API int cwGetProductByte(unsigned char ucProduct, unsigned char ucByte[]);


///@brief
///获取key剩余的有效期
///@param
///iType-有效期类型，1表示剩下有效小时数
///@return
///-1: 没有检测到key， -2: 检测到永久key, 0: key已过期， 其他：剩余的有效期
AUTHORIZE_API int cwGetValidity(int iType);


	/**
	 * 功能：获取key安装时间和有效时间，该接口获取有效时间最久的一个key，并不会考虑是什么类型的key
	 * 输入：
	        无
	 * 输出：
	 *      lStartTime          - key安装时间的时间戳，只有为限制时间的key，lStartTime和iValidTime才有值
	 *      iValidTime          - key有效时间天数
	 * 返回值：
	 *		int                 - -4：未检测到key；-3：key时间信息错误；-2：key已过期；-1：永久key，不限时间；其他(>0)：剩余有效小时数
	 */
AUTHORIZE_API int cwGetHaspKeyValidTime(long long *lStartTime, int *iValidTime);


	/**
	 * 功能：获取当前key和所有key的详细信息
	 * 输入：
	        iCurKeyBuf          - 获取当前key信息的buf长度
			iAllKeyBuf          - 获取所有key信息的buf长度
	 * 输出：
	 *      szCurKeyInfo        - 当前key的信息，包括类型，时间
	 *      szAllKeyInfo        - 当前机器装过的所有key信息
	 * 返回值：
	 *		int                 - 0：成功；-1：iCurKeyBuf长度不够；-2：iAllKeyBuf长度不够
	 */
AUTHORIZE_API int cwGetHaspKeyInfo(char szCurKeyInfo[], int iCurKeyBuf, char szAllKeyInfo[], int iAllKeyBuf);


    /**
	 * 功能：采集设备文件
	 * 输入：
	        iFpBuf              - 获取设备文件信息的buf长度
	 * 输出：
	 *      szFpFileData        - 设备文件数据
	 * 返回值：
	 *		int                 - 0：成功；其他：错误码
	 */
AUTHORIZE_API int cwGetFingerPrint(char szFpFileData[], int iFpBuf);


    /**
	 * 功能：采集升级文件
	 * 输入：
	        iUpdateBuf          - 获取升级文件信息的buf长度
	 * 输出：
	 *      szUpdateFileData    - 升级文件数据
	 * 返回值：
	 *		int                 - 0：成功；其他：错误码
	 */
AUTHORIZE_API int cwGetUpdateFile(char szUpdateFileData[], int iUpdateBuf);


    /**
	 * 功能：安装授权文件
	 * 输入：
	        sV2CFile            - 授权文件数据
	 * 输出：
	 *      无
	 * 返回值：
	 *		int                 - 0：成功；其他：错误码
	 */
AUTHORIZE_API int cwInstallFile(const char* sV2CFile);


#if defined (__cplusplus)
}
#endif

#endif

