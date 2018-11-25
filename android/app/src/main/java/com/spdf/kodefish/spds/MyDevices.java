package com.spdf.kodefish.spds;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ArrayAdapter;
import android.widget.ListView;
import android.widget.AdapterView.OnItemLongClickListener;
import android.support.v7.app.AlertDialog;
import android.content.DialogInterface;


import com.spdf.kodefish.spds.models.Devices;

import java.util.ArrayList;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class MyDevices extends AppCompatActivity {

    String token = "";

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_my_devices);
        Bundle extras = getIntent().getExtras();
        if (extras != null) {
            token = extras.getString("token");
        }

        // Call API to get my devices
        final SPDSBasicService restService = new SPDSBasicService();
        restService.spdsApi.getDevices(token).enqueue(new Callback<Devices>() {
            @Override
            public void onResponse(Call<Devices> call, Response<Devices> response) {
                // Populate list
                if (response.isSuccessful()) {
                    final ArrayList<String> otherUsers = new ArrayList<>(response.body().getDevices());
                    final ArrayAdapter<String> itemsAdapter = new ArrayAdapter<>(MyDevices.this, android.R.layout.simple_list_item_1, otherUsers);
                    final ListView listView = findViewById(R.id.list_view);
                    listView.setAdapter(itemsAdapter);
                    listView.setOnItemLongClickListener(new OnItemLongClickListener() {
                        @Override
                        public boolean onItemLongClick(AdapterView<?> parent, View view, final int position, long id) {
                            AlertDialog.Builder builder = new AlertDialog.Builder(MyDevices.this);
                            // Add the buttons
                            builder.setTitle("Remove Device");
                            builder.setMessage("Are you sure you want to remove this device?");
                            builder.setPositiveButton("Ok", new DialogInterface.OnClickListener() {
                                public void onClick(DialogInterface dialog, int id) {
                                    // User clicked OK button
                                    ArrayList<String> rm = new ArrayList<>();
                                    rm.add(otherUsers.get(position));
                                    Devices devicesToRm = new Devices();
                                    devicesToRm.setDevices(rm);
                                    restService.spdsApi.deleteDevices(devicesToRm, token).enqueue(new Callback<Devices>() {
                                        @Override
                                        public void onResponse(Call<Devices> call, Response<Devices> response) {
                                            if (response.isSuccessful()) {
                                                otherUsers.clear();
                                                if (response.body() != null)
                                                    otherUsers.addAll(response.body().getDevices());
                                                itemsAdapter.notifyDataSetChanged();
                                            }
                                        }

                                        @Override
                                        public void onFailure(Call<Devices> call, Throwable t) {

                                        }
                                    });
                                }
                            });
                            builder.setNegativeButton("Cancel", new DialogInterface.OnClickListener() {
                                public void onClick(DialogInterface dialog, int id) {
                                    // User cancelled the dialog
                                    dialog.dismiss();
                                }
                            });
                            // Set other dialog properties

                            // Create the AlertDialog
                            AlertDialog dialog = builder.create();
                            dialog.show();
                            return false;
                        }

                    });
                }

            }

            @Override
            public void onFailure(Call<Devices> call, Throwable t) {

            }
        });
    }
}
