package com.example.paybill.controller;

import com.example.paybill.dto.PayBillRequestDto;
import com.example.paybill.dto.PayBillResponseDto;
import com.example.paybill.service.PayBillService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequiredArgsConstructor
@Slf4j
public class PayBill {
    private final PayBillService payBillService;

    @PostMapping(value = "/api/pay-bill")
    public PayBillResponseDto payBill(@RequestBody PayBillRequestDto payBillRequest) {

        log.info("Pay Bill Request is {}", payBillRequest);
        return payBillService.payBill(payBillRequest);
    }
}
