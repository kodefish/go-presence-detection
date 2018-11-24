use spds;

var bulk = db.users.initializeUnorderedBulkOp();
bulk.insert({
    _id: 1, 
    username: "kodefish",
    devices: [
        {
            _id: "1",
            mac_address: "lulz_mac_addr",
            ip_address: "420.420.420.1",
            device_name: "OnePlus 3T",
            device_type: "0",
            mobile: true
        },
        {
            _id: "2",
            mac_address: "lulz_mac_addr",
            ip_address: "420.420.420.2",
            device_name: "OnePlus 3T2",
            device_type: "3",
            mobile: false
        },
    ]
});
bulk.execute();