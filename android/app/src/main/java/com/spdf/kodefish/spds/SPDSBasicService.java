package com.spdf.kodefish.spds;

import okhttp3.HttpUrl;
import okhttp3.OkHttpClient;
import okhttp3.logging.HttpLoggingInterceptor;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;

public class SPDSBasicService {
    public SPDSRestService spdsApi;

    SPDSBasicService() {
        HttpLoggingInterceptor interceptor = new HttpLoggingInterceptor();
        interceptor.setLevel(HttpLoggingInterceptor.Level.BODY);
        OkHttpClient client = new OkHttpClient.Builder().addInterceptor(interceptor).build();
        HttpUrl httpUrl = new HttpUrl.Builder().scheme("http").host("192.33.206.40").port(4200).build();
        Retrofit retrofit = new Retrofit.Builder()
                .baseUrl(httpUrl)
                .client(client)
                .addConverterFactory(GsonConverterFactory.create())
                .build();
        spdsApi = retrofit.create(SPDSRestService.class);
    }
}
