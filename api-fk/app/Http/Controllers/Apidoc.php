<?php

/**
 * @apiDefine BaseApiDefined
 * @apiHeader {String} ADMIN-TOKEN Authorization token.
 */

/**
 * @api {get} /x/x Common Header
 * @apiGroup A-CommonDoc
 * @apiSampleRequest off
 * @apiHeader {String} ADMIN-TOKEN
 * @apiDescription
 *  common header
 */

/**
 * @api {get} /x/xx response structure
 * @apiGroup A-CommonDoc
 * @apiSampleRequest off
 *
 * @apiDescription response json
 * @apiSuccessExample {json} Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "data": "success data",
 *       "code": "error code",
 *       "msg": "message"
 *     }
 */

/**
 * @api {get} /x/xxx Global Error Code
 * @apiGroup A-CommonDoc
 * @apiSampleRequest off
 * @apiError (Sys) -1 System Error.
 * @apiError (Sys) 1 .Validation Error.
 */
