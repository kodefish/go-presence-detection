package com.spdf.kodefish.spds;

import com.spdf.kodefish.spds.models.Devices;
import com.spdf.kodefish.spds.models.JWTToken;
import com.spdf.kodefish.spds.models.User;

import java.util.List;

import okhttp3.ResponseBody;
import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.GET;
import retrofit2.http.Header;
import retrofit2.http.POST;

public interface SPDSRestService {
    @POST("/register")
    Call<JWTToken> register(@Body User user);

    @POST("/get-token")
    Call<JWTToken> getToken(@Body User user);

    @GET("/devices")
    Call<Devices> getDevices(@Header("Authorization") String token);

    @GET("connected-users")
    Call<List<String>> getConnectedUsers(@Header("Authorization") String token);

    @GET("/add-device")
    Call<ResponseBody> addDevice(@Header("Authorization") String token);

    @POST("/delete-device")
    Call<Devices> deleteDevices(@Body Devices devices, @Header("Authorization") String token);



}
