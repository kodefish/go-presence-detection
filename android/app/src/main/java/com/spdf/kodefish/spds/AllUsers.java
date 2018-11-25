package com.spdf.kodefish.spds;

import android.content.Intent;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.Menu;
import android.view.MenuInflater;
import android.view.MenuItem;
import android.widget.ListView;
import android.widget.ArrayAdapter;
import android.widget.Toast;

import java.util.ArrayList;
import java.util.List;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class AllUsers extends AppCompatActivity {

    String token = "";

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_all_users);
        Bundle extras = getIntent().getExtras();
        if (extras != null) {
            token = "Bearer " + extras.getString("token");
        }
        downloadAllUsers();


    }

    private void downloadAllUsers() {
        // Init retrofit
        SPDSBasicService restService = new SPDSBasicService();

        // Download other users
        restService.spdsApi.getConnectedUsers(token).enqueue(new Callback<List<String>>() {
            @Override
            public void onResponse(Call<List<String>> call, Response<List<String>> response) {
                if (response.isSuccessful() && response.body() != null) {
                    ArrayList<String> otherUsers = new ArrayList<>(response.body());
                    ArrayAdapter<String> itemsAdapter = new ArrayAdapter<String>(AllUsers.this, android.R.layout.simple_list_item_1, otherUsers);
                    ListView listView = (ListView) findViewById(R.id.list_view);
                    listView.setAdapter(itemsAdapter);
                }
            }

            @Override
            public void onFailure(Call<List<String>> call, Throwable t) {

            }
        });
    }

    @Override
    protected void onResume() {
        super.onResume();
        downloadAllUsers();
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        MenuInflater inflater = getMenuInflater();
        inflater.inflate(R.menu.option_menu, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        switch (item.getItemId()) {
            case R.id.refresh:
                downloadAllUsers();
            case R.id.add_this_device:
                addMyDevice();
                return true;
            case R.id.my_devices:
                openMyDevices();
                return true;
                default:
                    return super.onOptionsItemSelected(item);
        }
    }

    private void openMyDevices() {
        Intent myDevices = new Intent(this, MyDevices.class);
        myDevices.putExtra("token", token);
        startActivity(myDevices);
    }

    private void addMyDevice() {
        SPDSBasicService restService = new SPDSBasicService();
        restService.spdsApi.addDevice(token).enqueue(new Callback() {
            @Override
            public void onResponse(Call call, Response response) {
                if (response.isSuccessful()) {
                    Toast.makeText(AllUsers.this, "Your device has been added!", Toast.LENGTH_SHORT);
                } else {
                    Toast.makeText(AllUsers.this, "Unfortunately, something went wrong", Toast.LENGTH_SHORT);
                }
            }

            @Override
            public void onFailure(Call call, Throwable t) {
                Toast.makeText(AllUsers.this, "Unfortunately, something went wrong", Toast.LENGTH_SHORT);
            }
        });
    }
}
