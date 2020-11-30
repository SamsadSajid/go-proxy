package com.example.paybill.dto;

import lombok.Builder;
import lombok.Data;

import java.util.Date;

@Data
@Builder
public class PayBillResponseDto {
    private String code;
    private String message;
    private Date updateTime;
}
