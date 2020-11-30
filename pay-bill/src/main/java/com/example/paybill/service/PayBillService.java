package com.example.paybill.service;

import com.example.paybill.dto.PayBillRequestDto;
import com.example.paybill.dto.PayBillResponseDto;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.concurrent.*;

@Service
@Slf4j
public class PayBillService {

    public PayBillResponseDto payBill(PayBillRequestDto payBillRequest) {

        Callable<PayBillResponseDto> payBillTask = () -> {
            log.info("Accessing DB...");

            log.info("Calling service A...");

            log.info("Calling service B...");


            return PayBillResponseDto.builder()
                    .code("1234")
                    .message("Bill pay is successful")
                    .updateTime(new Date())
                    .build();
        };

        ScheduledExecutorService scheduledExecutorService = Executors.newScheduledThreadPool(1);

        ScheduledFuture<PayBillResponseDto> payBillResponseScheduledFuture =
                scheduledExecutorService.schedule(payBillTask, 1500, TimeUnit.MILLISECONDS);

        scheduledExecutorService.shutdown();

        try {
            return payBillResponseScheduledFuture.get();
        } catch (Exception e) {
            e.printStackTrace();
            throw new RuntimeException();
        }
    }
}
