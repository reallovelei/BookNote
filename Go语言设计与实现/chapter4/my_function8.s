	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 10, 15	sdk_version 10, 15
	.globl	_my_function            ## -- Begin function my_function
	.p2align	4, 0x90
_my_function:                           ## @my_function
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	movl	24(%rbp), %eax
	movl	16(%rbp), %r10d
	movl	%edi, -4(%rbp)
	movl	%esi, -8(%rbp)
	movl	%edx, -12(%rbp)
	movl	%ecx, -16(%rbp)
	movl	%r8d, -20(%rbp)
	movl	%r9d, -24(%rbp)
	movl	-4(%rbp), %ecx
	addl	-8(%rbp), %ecx
	addl	-12(%rbp), %ecx
	addl	-16(%rbp), %ecx
	addl	-20(%rbp), %ecx
	addl	-24(%rbp), %ecx
	addl	16(%rbp), %ecx
	addl	24(%rbp), %ecx
	movl	%eax, -28(%rbp)         ## 4-byte Spill
	movl	%ecx, %eax
	movl	%r10d, -32(%rbp)        ## 4-byte Spill
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function
	.globl	_main                   ## -- Begin function main
	.p2align	4, 0x90
_main:                                  ## @main
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	subq	$32, %rsp
	movl	$1, %edi
	movl	$2, %esi
	movl	$3, %edx
	movl	$4, %ecx
	movl	$5, %r8d
	movl	$6, %r9d
	movl	$7, (%rsp)
	movl	$16, 8(%rsp)
	callq	_my_function
	xorl	%ecx, %ecx
	movl	%eax, -4(%rbp)
	movl	%ecx, %eax
	addq	$32, %rsp
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function

.subsections_via_symbols
