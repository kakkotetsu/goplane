// Copyright (C) 2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

//
// from include/linux/rtnetlink.h
//
// /* RTnetlink multicast groups - backwards compatibility for userspace */
// #define RTMGRP_LINK		1
// #define RTMGRP_NOTIFY	2
// #define RTMGRP_NEIGH		4
// #define RTMGRP_TC		8
//
// #define RTMGRP_IPV4_IFADDR	0x10
// #define RTMGRP_IPV4_MROUTE	0x20
// #define RTMGRP_IPV4_ROUTE	0x40
// #define RTMGRP_IPV4_RULE	0x80
//
// #define RTMGRP_IPV6_IFADDR	0x100
// #define RTMGRP_IPV6_MROUTE	0x200
// #define RTMGRP_IPV6_ROUTE	0x400
// #define RTMGRP_IPV6_IFINFO	0x800
//
// #define RTMGRP_DECnet_IFADDR    0x1000
// #define RTMGRP_DECnet_ROUTE     0x4000
//
// #define RTMGRP_IPV6_PREFIX	0x20000

type RTMGRP_TYPE uint

const (
	_ RTMGRP_TYPE = iota
	RTMGRP_LINK
	RTMGRP_NOTIFY
	RTMGRP_NEIGH
	RTMGRP_TC
	RTMGRP_IPV4_IFADDR
	RTMGRP_IPV4_MROUTE
	RTMGRP_IPV4_ROUTE
	RTMGRP_IPV4_RULE
	RTMGRP_IPV6_IFADDR
	RTMGRP_IPV6_MROUTE
	RTMGRP_IPV6_ROUTE
	RTMGRP_IPV6_IFINFO
	RTMGRP_DECnet_IFADDR
	_
	RTMGRP_DECnet_ROUTE
	_
	RTMGRP_IPV6_PREFIX
)

type RTM_TYPE uint

const (
	RTM_NEWLINK RTM_TYPE = 16
	RTM_DELLINK RTM_TYPE = 17
	RTM_GETLINK RTM_TYPE = 18
	RTM_SETLINK RTM_TYPE = 19

	RTM_NEWADDR RTM_TYPE = 20
	RTM_DELADDR RTM_TYPE = 21
	RTM_GETADDR RTM_TYPE = 22

	RTM_NEWROUTE RTM_TYPE = 24
	RTM_DELROUTE RTM_TYPE = 25
	RTM_GETROUTE RTM_TYPE = 26

	RTM_NEWNEIGH RTM_TYPE = 28
	RTM_DELNEIGH RTM_TYPE = 29
	RTM_GETNEIGH RTM_TYPE = 30

	RTM_NEWRULE RTM_TYPE = 32
	RTM_DELRULE RTM_TYPE = 33
	RTM_GETRULE RTM_TYPE = 34

	RTM_NEWQDISC RTM_TYPE = 36
	RTM_DELQDISC RTM_TYPE = 37
	RTM_GETQDISC RTM_TYPE = 38

	RTM_NEWTCLASS RTM_TYPE = 40
	RTM_DELTCLASS RTM_TYPE = 41
	RTM_GETTCLASS RTM_TYPE = 42

	RTM_NEWTFILTER RTM_TYPE = 44
	RTM_DELTFILTER RTM_TYPE = 45
	RTM_GETTFILTER RTM_TYPE = 46

	RTM_NEWACTION RTM_TYPE = 48
	RTM_DELACTION RTM_TYPE = 49
	RTM_GETACTION RTM_TYPE = 50

	RTM_NEWPREFIX RTM_TYPE = 52

	RTM_GETMULTICAST RTM_TYPE = 58

	RTM_GETANYCAST RTM_TYPE = 62

	RTM_NEWNEIGHTBL RTM_TYPE = 64
	RTM_GETNEIGHTBL RTM_TYPE = 66
	RTM_SETNEIGHTBL RTM_TYPE = 67

	RTM_NEWNDUSEROPT RTM_TYPE = 68

	RTM_NEWADDRLABEL RTM_TYPE = 72
	RTM_DELADDRLABEL RTM_TYPE = 73
	RTM_GETADDRLABEL RTM_TYPE = 74

	RTM_GETDCB RTM_TYPE = 78
	RTM_SETDCB RTM_TYPE = 79

	RTM_NEWNETCONF RTM_TYPE = 80
	RTM_GETNETCONF RTM_TYPE = 82

	RTM_NEWMDB RTM_TYPE = 84
	RTM_DELMDB RTM_TYPE = 85
	RTM_GETMDB RTM_TYPE = 86
)

type NDA_TYPE uint

const (
	NDA_UNSPEC NDA_TYPE = iota
	NDA_DST
	NDA_LLADDR
	NDA_CACHEINFO
	NDA_PROBES
	NDA_VLAN
	NDA_PORT
	NDA_VNI
	NDA_IFINDEX
	NDA_MAX = NDA_IFINDEX
)

// Neighbor Cache Entry States.
type NUD_TYPE uint

const (
	NUD_NONE       NUD_TYPE = 0x00
	NUD_INCOMPLETE NUD_TYPE = 0x01
	NUD_REACHABLE  NUD_TYPE = 0x02
	NUD_STALE      NUD_TYPE = 0x04
	NUD_DELAY      NUD_TYPE = 0x08
	NUD_PROBE      NUD_TYPE = 0x10
	NUD_FAILED     NUD_TYPE = 0x20
	NUD_NOARP      NUD_TYPE = 0x40
	NUD_PERMANENT  NUD_TYPE = 0x80
)

// Neighbor Flags
type NTF_TYPE uint

const (
	NTF_USE    NTF_TYPE = 0x01
	NTF_SELF   NTF_TYPE = 0x02
	NTF_MASTER NTF_TYPE = 0x04
	NTF_PROXY  NTF_TYPE = 0x08
	NTF_ROUTER NTF_TYPE = 0x80
)
