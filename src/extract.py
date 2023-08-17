import os
import json
from facebook_business.api import FacebookAdsApi
from facebook_business.adobjects.adaccount import AdAccount
from facebook_business.adobjects.campaign import Campaign
from facebook_business.adobjects.adset import AdSet
from facebook_business.adobjects.ad import Ad
from facebook_business.exceptions import FacebookRequestError
from google.cloud import storage
from dotenv import load_dotenv
import json
import backoff
import logging
from time import sleep
load_dotenv() # For local testing

# GLOBAL VARIABLES
ACCESS_TOKEN=os.environ['ACCESS_TOKEN']
ACCOUNT_ID = os.environ['AD_ACCOUNT_ID']
STATE_BUCKET = os.environ['GCS_STATE_BUCKET']
DEST_BUCKET = os.environ['GCS_DEST_BUCKET']

# GOOGLE CLOUD STORAGE
gcs = None
def get_gcs_client():
    global gcs
    if gcs == None:
        gcs = storage.Client()
    return gcs

# Enabling logging for backoff
logging.getLogger('backoff').addHandler(logging.StreamHandler())


# FACEBOOK BUSINESS SDK FUNCTIONS
@backoff.on_exception(backoff.expo, FacebookRequestError, max_time=120)
def _get_cursor(func, fields=None, params=None):
    return func(fields=fields, params=params)

@backoff.on_exception(backoff.expo, FacebookRequestError, max_time=120)
def _get_value(cursor):
    return next(cursor)

def _get_collection(cursor):

    collection = []
    for node in cursor:
        collection.append(node)
    return collection


def get_all_campaigns(state = None):
    
    after = None
    if state != None:
        after = state['after']

    fields = [
      'name',
      'objective',
    ]
    params = {
            'effective_status': ['ACTIVE','PAUSED'], 'limit': 500, 'after': after
    }

    # Create the cursor
    cursor = _get_cursor(
        AdAccount(ACCOUNT_ID).get_campaigns,
        fields,
        params
    )
    # Get get_campaigns
    return _get_collection(cursor)


def get_ad_sets():
    
    # Define fields and params
    fields = [
      'adset_id'
    ]
    params = {'limit': 500}

    cursor = _get_cursor(
            AdAccount(ACCOUNT_ID).get_ad_sets,
            fields,
            params
    )

    return _get_collection(cursor)


def get_all_ads():

    # Define fields and params
    # TODO
    params = {'limit': 500}

    # Create cursor
    cursor = _get_cursor(
            AdAccount(ACCOUNT_ID).get_ads,
            params = params
    )
    
    return _get_collection(cursor)
    

def get_ads_insights():
    fields = [
        "account_id",
        "account_name",
        "ad_id",
        "ad_name",
        "adset_id",
        "adset_name",
        "campaign_id",
        "campaign_name",
        "clicks",
        "conversions",
        "cost_per_ad_click",
        "cpc",
        "cpm",
        "cpp",
        "created_time",
        "ctr",
        "date_start",
        "date_stop",
        "frequency",
        "full_view_impressions",
        "full_view_reach",
        "impressions",
        "reach",
        "social_spend",
        "spend",
        "updated_time"
    ]
    params = {
        'level': 'ad',
        'date_preset': 'maximum',
        'limit': 500
    }

    # Get cursor
    cursor = _get_cursor(
        AdAccount(ID).get_insights,
        fields=fields,
        params=params,
    )
    return _get_collection(cursor)


def main():
    
    FacebookAdsApi.init(access_token=ACCESS_TOKEN)

    c = get_all_campaigns()
    print(len(c))
    print(c[-1])
    s = get_ad_sets()
    print(len(s))
    print(s[-1])
    a = get_all_ads()
    print(len(a))
    print(a[-1])

if __name__ == "__main__":
    main()
