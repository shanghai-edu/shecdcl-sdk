'use strict'
var log4js = require('log4js')
log4js.configure('/root/nodejs/log4js.json')
const config = require('../example/config')


var assert = require('assert');
var GW = require('../shecgw');

const XM = process.env.SHECGW_TEST_XM || '王*平';
const ZJHM = process.env.SHECGW_TEST_ZJHM || '身份证号码HERE';

describe('shecgw', function (){
    var gw ;
    before(function(){
        gw = new GW(config);
    })
    describe('#getAccessToken()', function(){
        it('should get access token', async function(){
            let token = await gw.getAccessToken();
            console.log(token);
            assert.ok(token);
        })
    });

    describe('#GetSjSsmjm()', function(){
        it('should get 00', async function(){
            let ret = await gw.GetSjSsmjm(XM, ZJHM);
            console.log(ret);  
            assert.equal(ret.code, '0');
        })
    });

    describe('#GetGjXgfyHsjcsjfwjk', function(){
        it('shoud get none', async function(){
            let ret = await gw.GetGjXgfyHsjcsjfwjk(XM, ZJHM);

            assert.equal(ret.code, 1);
        })
    });

    describe('#GetGjXgymjzxx', function(){
        it('should get 3 times', async function(){
            let ret = await gw.GetGjXgymjzxx(XM, ZJHM);

            assert.equal(ret.code, '0');
        })
    })

    describe('#getHscyxxcx',function(){
        it('should get hscy inforation', async function(){
            let ret = await gw.GetHscyxxcx(ZJHM);

            assert.eequal(ret.code, '0');
        })
    })

})
