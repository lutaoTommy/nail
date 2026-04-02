package handler

import (
    "regexp"
)

/*设备信息*/
type Device struct {
    Id           string  `gorm:"primaryKey;column:id;size:20;" json:"id"`
    Mac          string  `gorm:"column:mac;size:20;" json:"mac"`
    Bat          string  `gorm:"column:bat;size:10" json:"bat"`
    Name         string  `gorm:"column:name;size:50" json:"name"`
    Rssi         int     `gorm:"column:rssi;type:int" json:"rssi"`
    Count        int     `gorm:"column:count;type:int" json:"count"`
    Time         string  `gorm:"column:time;size:20" json:"time"`
    Version      string  `gorm:"column:version;size:30" json:"version"`
    UserId       string  `gorm:"column:user_id;size:20" json:"user_id"`
}

/*设备上报*/
type DeviceUpload struct {
    Token        string   `json:"token"`
    Devices      []Device `json:"devices"`
}

/*字段检查*/
func (device Device) checkMac() error {
    var err error
    rgx := regexp.MustCompile(MacReg)
    if device.Mac == "" {
        err = newError(400, "E_NO_MAC")
    } else if !rgx.MatchString(device.Mac) {
        err = newError(400, "E_INVALID_MAC")
    } else if len(device.Mac) > 20 {
        err = newError(400, "E_TOO_LONG")
    }
    return err
}

/*字段检查*/
func (device Device) checkParams() error {
    var err error
    if len(device.Bat) > 10  {
        err = newError(400, "E_TOO_LONG")
    } else if len(device.Version) > 30 {
        err = newError(400, "E_TOO_LONG")
    } else if len(device.Name) > 50 {
        err = newError(400, "E_TOO_LONG")
    }
    return err
}


/*换色信息*/
type PostInfo struct {
    Id           string       `gorm:"primaryKey;column:id;size:20" json:"id"`
    Mac          string       `gorm:"column:mac;size:20;" json:"mac"`
    Time         string       `gorm:"column:time;size:20" json:"time"`
    Token        string       `gorm:"-" json:"token"`
    UserId       string       `gorm:"column:user_id;size:20" json:"user_id"`
    DeviceName   string       `gorm:"column:device_name;size:50" json:"device_name"`
    Colors       []PostColor  `gorm:"foreignKey:PostId" json:"colors"`
}

/*换色信息*/
type PostInfoOut struct {
    Id           string       `json:"id"`
    Mac          string       `json:"mac"`
    Time         string       `json:"time"`
    DeviceName   string       `json:"device_name"`
    Colors       []PostColor  `gorm:"foreignKey:PostId" json:"colors"`
}

/*换色信息*/
type PostColor struct {
    No         int     `gorm:"column:no;type:int" json:"no"`
    Id         string  `gorm:"column:id;size:5;" json:"id"`
    R          int     `gorm:"column:r;type:int" json:"r"`
    G          int     `gorm:"column:g;type:int" json:"g"`
    B          int     `gorm:"column:b;type:int" json:"b"`
    X          int     `gorm:"column:x;type:int" json:"x"`
    Y          int     `gorm:"column:y;type:int" json:"y"`
    Desc       string  `gorm:"column:desc;size:10" json:"desc"`
    Name       string  `gorm:"column:name;size:30" json:"name"`
    Color      string  `gorm:"column:color;size:30" json:"color"`
    ColorId    string  `gorm:"primaryKey;column:color_id;size:30" json:"color_id"`
    PostId     string  `gorm:"column:post_id;size:20;index" json:"post_id"`
}

/*设备别名*/
type DeviceName struct {
    Mac          string  `gorm:"column:mac;size:20;" json:"mac"`
    Time         string  `gorm:"column:time;size:20" json:"time"`
    Name         string  `gorm:"column:name;size:50" json:"name"`
}

