section .data
    arr     dq 81, 17, 36, 25, 36, 99, 144
    arr_len dq 7
    fmt     db "Result: %d", 10, 0

section .text
    global main
    extern printf

main:
    push rbp
    mov rbp, rsp
    and rsp, -16
    sub rsp, 32

    xor rdx, rdx
    mov dword [rbp-4], 0

    lea r8, [rel arr]

next:
    cmp rdx, [rel arr_len]
    je end
    mov rax, [r8 + rdx*8]
    inc rdx

    xor rbx, rbx

check_square:
    inc rbx
    mov rcx, rbx
    imul rcx, rcx
    cmp rcx, rax
    jl check_square
    jg next
    inc dword [rbp-4]
    jmp next

end:
    lea rdi, [rel fmt]
    mov esi, [rbp-4]
    xor eax, eax
    call printf

    xor eax, eax
    leave
    ret
