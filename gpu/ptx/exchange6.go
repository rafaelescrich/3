package ptx

const EXCHANGE6 = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Wed Aug  1 02:51:19 2012 (1343782279)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_00001b24_00000000-9_exchange6.cpp3.i"
	.file	2 "/home/arne/src/nimble-cube/gpu/ptx/exchange6.cu"

.visible .entry _Z9exchange6PfS_fffiiiiii(
	.param .u64 _Z9exchange6PfS_fffiiiiii_param_0,
	.param .u64 _Z9exchange6PfS_fffiiiiii_param_1,
	.param .f32 _Z9exchange6PfS_fffiiiiii_param_2,
	.param .f32 _Z9exchange6PfS_fffiiiiii_param_3,
	.param .f32 _Z9exchange6PfS_fffiiiiii_param_4,
	.param .u32 _Z9exchange6PfS_fffiiiiii_param_5,
	.param .u32 _Z9exchange6PfS_fffiiiiii_param_6,
	.param .u32 _Z9exchange6PfS_fffiiiiii_param_7,
	.param .u32 _Z9exchange6PfS_fffiiiiii_param_8,
	.param .u32 _Z9exchange6PfS_fffiiiiii_param_9,
	.param .u32 _Z9exchange6PfS_fffiiiiii_param_10
)
{
	.reg .pred 	%p<17>;
	.reg .s32 	%r<89>;
	.reg .f32 	%f<22>;
	.reg .s64 	%rd<24>;


	ld.param.u64 	%rd2, [_Z9exchange6PfS_fffiiiiii_param_0];
	ld.param.u64 	%rd3, [_Z9exchange6PfS_fffiiiiii_param_1];
	ld.param.f32 	%f3, [_Z9exchange6PfS_fffiiiiii_param_3];
	ld.param.f32 	%f4, [_Z9exchange6PfS_fffiiiiii_param_4];
	ld.param.u32 	%r43, [_Z9exchange6PfS_fffiiiiii_param_5];
	ld.param.u32 	%r44, [_Z9exchange6PfS_fffiiiiii_param_6];
	ld.param.u32 	%r45, [_Z9exchange6PfS_fffiiiiii_param_7];
	ld.param.u32 	%r46, [_Z9exchange6PfS_fffiiiiii_param_8];
	ld.param.u32 	%r47, [_Z9exchange6PfS_fffiiiiii_param_9];
	ld.param.u32 	%r48, [_Z9exchange6PfS_fffiiiiii_param_10];
	.loc 2 6 1
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	.loc 2 7 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 9 1
	setp.lt.s32 	%p1, %r8, %r48;
	setp.lt.s32 	%p2, %r4, %r47;
	and.pred  	%p3, %p2, %p1;
	.loc 2 13 1
	setp.gt.s32 	%p4, %r46, 0;
	.loc 2 9 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_9;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 36 1
	mad.lo.s32 	%r9, %r4, %r48, %r8;
	.loc 2 25 1
	add.s32 	%r50, %r46, -1;
	mad.lo.s32 	%r51, %r50, %r47, %r4;
	mad.lo.s32 	%r10, %r51, %r48, %r8;
	.loc 2 13 1
	add.s32 	%r52, %r47, -1;
	mad.lo.s32 	%r85, %r48, %r52, %r8;
	mul.lo.s32 	%r12, %r48, %r47;
	add.s32 	%r54, %r4, 1;
	mad.lo.s32 	%r84, %r48, %r54, %r8;
	add.s32 	%r55, %r4, -1;
	mad.lo.s32 	%r83, %r48, %r55, %r8;
	mad.lo.s32 	%r82, %r48, %r54, -1;
	mul.lo.s32 	%r81, %r48, %r4;
	mad.lo.s32 	%r80, %r48, %r4, %r8;
	sub.s32 	%r56, %r4, %r47;
	mul.lo.s32 	%r19, %r48, %r56;
	add.s32 	%r57, %r3, %r47;
	mad.lo.s32 	%r58, %r2, %r1, %r57;
	mul.lo.s32 	%r20, %r48, %r58;
	mov.u32 	%r86, 0;
	cvta.to.global.u64 	%rd21, %rd2;
	mov.u32 	%r79, %r8;

BB0_2:
	.loc 2 32 1
	mov.u32 	%r21, %r79;
	cvta.to.global.u64 	%rd4, %rd3;
	.loc 2 16 1
	cvt.s64.s32 	%rd1, %r80;
	mul.wide.s32 	%rd5, %r80, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.f32 	%f1, [%rd6];
	.loc 2 21 1
	setp.gt.s32 	%p6, %r86, 0;
	@%p6 bra 	BB0_4;

	setp.eq.s32 	%p7, %r43, 0;
	.loc 2 24 1
	selp.b32 	%r87, %r80, %r10, %p7;
	bra.uni 	BB0_5;

BB0_4:
	.loc 2 22 1
	add.s32 	%r87, %r19, %r21;

BB0_5:
	.loc 2 30 1
	mul.wide.s32 	%rd8, %r87, 4;
	add.s64 	%rd9, %rd4, %rd8;
	ld.global.f32 	%f2, [%rd9];
	.loc 2 32 1
	add.s32 	%r86, %r86, 1;
	setp.lt.s32 	%p8, %r86, %r46;
	@%p8 bra 	BB0_7;

	setp.eq.s32 	%p9, %r43, 0;
	.loc 2 35 1
	selp.b32 	%r88, %r80, %r9, %p9;
	bra.uni 	BB0_8;

BB0_7:
	.loc 2 33 1
	add.s32 	%r88, %r20, %r21;

BB0_8:
	.loc 2 82 1
	add.s32 	%r61, %r8, 1;
	setp.lt.s32 	%p10, %r61, %r48;
	setp.lt.s32 	%p11, %r54, %r47;
	.loc 2 41 1
	mul.wide.s32 	%rd11, %r88, 4;
	add.s64 	%rd12, %rd4, %rd11;
	ld.global.f32 	%f5, [%rd12];
	.loc 2 43 1
	sub.ftz.f32 	%f6, %f5, %f1;
	sub.ftz.f32 	%f7, %f2, %f1;
	add.ftz.f32 	%f8, %f7, %f6;
	setp.eq.s32 	%p12, %r44, 0;
	.loc 2 49 1
	selp.b32 	%r64, %r80, %r85, %p12;
	setp.gt.s32 	%p13, %r4, 0;
	.loc 2 46 1
	selp.b32 	%r65, %r83, %r64, %p13;
	.loc 2 55 1
	mul.wide.s32 	%rd13, %r65, 4;
	add.s64 	%rd14, %rd4, %rd13;
	.loc 2 60 1
	selp.b32 	%r66, %r80, %r21, %p12;
	.loc 2 57 1
	selp.b32 	%r67, %r84, %r66, %p11;
	.loc 2 66 1
	mul.wide.s32 	%rd15, %r67, 4;
	add.s64 	%rd16, %rd4, %rd15;
	.loc 2 55 1
	ld.global.f32 	%f9, [%rd14];
	.loc 2 68 1
	sub.ftz.f32 	%f10, %f9, %f1;
	.loc 2 66 1
	ld.global.f32 	%f11, [%rd16];
	.loc 2 68 1
	sub.ftz.f32 	%f12, %f11, %f1;
	add.ftz.f32 	%f13, %f10, %f12;
	mul.ftz.f32 	%f14, %f13, %f3;
	fma.rn.ftz.f32 	%f15, %f8, %f4, %f14;
	setp.eq.s32 	%p14, %r45, 0;
	.loc 2 74 1
	selp.b32 	%r70, %r80, %r82, %p14;
	.loc 2 71 1
	add.s32 	%r71, %r80, -1;
	setp.gt.s32 	%p15, %r8, 0;
	.loc 2 71 1
	selp.b32 	%r72, %r71, %r70, %p15;
	.loc 2 80 1
	mul.wide.s32 	%rd17, %r72, 4;
	add.s64 	%rd18, %rd4, %rd17;
	.loc 2 85 1
	selp.b32 	%r73, %r80, %r81, %p14;
	.loc 2 82 1
	add.s32 	%r74, %r80, 1;
	selp.b32 	%r75, %r74, %r73, %p10;
	.loc 2 91 1
	mul.wide.s32 	%rd19, %r75, 4;
	add.s64 	%rd20, %rd4, %rd19;
	.loc 2 80 1
	ld.global.f32 	%f16, [%rd18];
	.loc 2 93 1
	sub.ftz.f32 	%f17, %f16, %f1;
	.loc 2 91 1
	ld.global.f32 	%f18, [%rd20];
	.loc 2 93 1
	sub.ftz.f32 	%f19, %f18, %f1;
	add.ftz.f32 	%f20, %f17, %f19;
	fma.rn.ftz.f32 	%f21, %f20, %f4, %f15;
	.loc 2 96 1
	shl.b64 	%rd22, %rd1, 2;
	add.s64 	%rd23, %rd21, %rd22;
	st.global.f32 	[%rd23], %f21;
	.loc 2 32 1
	add.s32 	%r85, %r85, %r12;
	add.s32 	%r84, %r84, %r12;
	add.s32 	%r83, %r83, %r12;
	add.s32 	%r82, %r82, %r12;
	add.s32 	%r81, %r81, %r12;
	add.s32 	%r80, %r80, %r12;
	add.s32 	%r42, %r21, %r12;
	.loc 2 13 1
	mov.u32 	%r79, %r42;
	@%p8 bra 	BB0_2;

BB0_9:
	.loc 2 98 2
	ret;
}


`
